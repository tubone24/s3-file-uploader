<template>
    <div id="list">
        <div class="select">
            <b-field>
                <b-select id="format-select" v-model="state.selected">
                    <option disabled value="">ファイル種類選択</option>
                    <option value="test">test</option>
                </b-select>
            </b-field>
            <div v-if="state.selected">
                <b-button id="search-file" type="is-primary" @click="updateFileList">検索</b-button>
            </div>
        </div>
        <div v-if="state.isLoading">
            <div v-show="state.isLoading" id="post-file-loader">
                <vue-loading type="bars" color="#0099e4" :size="{ width: '200px', height: '100px' }"></vue-loading>
                ダウンロード中...
            </div>
        </div>
        <div class="file-table">
            <b-table
               :data="state.uploadList"
               :paginated="true"
               :pagination-position="bottom"
               :default-sort="updated"
               :sticky-header="stickyHeaders"
            >
              <template slot-scope="uploadList">
                <b-table-column fileld="filename" label="ファイル名" width="120" sortable>
                  {{uploadList.row.name}}
                </b-table-column>
                <b-table-column fileld="updated" label="更新日時" width="120" sortable>
                  {{uploadList.row.lastModified}}
                </b-table-column>
                <b-table-column fileld="size" label="サイズ" width="100" sortable>
                  {{uploadList.row.size}}
                </b-table-column>
                <b-table-column fileld="download" label="-" width="100" sortable>
                  <button class="download-button" v-on:click="doDownload(uploadList.row.name)">DL </button>
                  <button class="delete-button" v-on:click="doComfirmDelete(uploadList.row.name)">DEL</button>
                </b-table-column>
              </template>
            </b-table>
<!--        <table>-->
<!--            <thead>-->
<!--            <tr>-->
<!--                <th class="filepath">ファイル名</th>-->
<!--                <th class="updated">更新日時</th>-->
<!--                <th class="size">サイズ</th>-->
<!--                <th class="download">-</th>-->
<!--            </tr>-->
<!--            </thead>-->
<!--            <tbody>-->
<!--            <tr v-for="item in state.uploadList">-->
<!--                <th>{{ item.name }}</th>-->
<!--                <th>{{ item.lastModified }}</th>-->
<!--                <th>{{ item.size }}</th>-->
<!--                <button class="download-button" v-on:click="doDownload(item.name)">DL</button>-->
<!--                <button class="delete-button" v-on:click="doComfirmDelete(item.name)">DEL</button>-->
<!--            </tr>-->
<!--            </tbody>-->
<!--        </table>-->
        </div>
        <v-dialog/>
    </div>
</template>

<script lang="ts">
    import {
        createComponent,
        reactive,
        onBeforeMount,
        SetupContext,
        onMounted,
        computed,
        ref
    } from '@vue/composition-api';
    import axios from 'axios';
    import { VueLoading } from 'vue-loading-template';
    import {FileNotFoundError} from "~/types/error";
    // data
    const state = reactive<{
        uploadList: Array<Map<string, string>>
        selected: string,
        isLoading: boolean,
        modalSelect: string
    }>({
        uploadList: [],
        selected: '',
        isLoading: false,
        modalSelect: ''
    });

    const updateFileList = async (): Promise<void> => {
        const res = await axios.get('/api/list', {
            params: {
                prefix: state.selected
            }
        });
        if (res.status === 200) {
            state.uploadList = res.data.fileList;
        }
        //console.log(JSON.stringify(state.uploadList));
    };

    const downloadFile = async (filePath: string): Promise<Blob> => {
        const res = await axios.get('/api/download', {
            params: {
                key: filePath
            }
        }
        ).catch((err) => {
                if (err.response.status === 404) {
                    throw new FileNotFoundError('FileFileNotFound');
                } else {
                    throw err;
                }
            },
        );
        return new Blob([res.data]);
    };
    type Props = {
        propHello: string;
    };
    export default createComponent({
        components: {
            VueLoading
        },
        props: {
            propHello: {
                type: String
            }
        },
        setup (props: Props, ctx:SetupContext) {
            // props
            const propsHello = props.propHello;
            const toast = ctx.root.$root.$toast;
            const modal = ctx.root.$root.$modal;
            const doDownload = async (filePath: string): Promise<void> => {
                state.isLoading = true;
                try{
                    const blob = await downloadFile(filePath);
                    const link = document.createElement('a');
                    link.href = window.URL.createObjectURL(blob);
                    const filename = filePath.match(".+/(.+?)([\?#;].*)?$")[1];
                    link.download = filename;
                    link.click();
                    state.isLoading = false;
                } catch (e) {
                    if (e instanceof FileNotFoundError) {
                        console.log(toast);
                        state.isLoading = false;
                        toast.global.nofileError();
                    } else {
                        state.isLoading = false;
                        toast.global.unknownError();
                    }
                }
            };
            const doDeleteFile = async (): Promise<void> => {
                modal.hide('dialog');
                let statusCode = 404;
                const res = await axios.post('/api/delete', {
                    key: state.modalSelect
                }).then((response) => {
                    statusCode = response.status;
                    toast.global.allSuccessDelete();
                }).catch((err) => {
                    statusCode = err.response.status;
                    toast.global.unknownError();
                });
                await updateFileList();
            };
            const doComfirmDelete = (filePath: string) => {
                state.modalSelect = filePath;
                modal.show('dialog', {
                    title: '確認',
                    text: '本当に削除していいですか？',
                    buttons: [
                        {
                            title: 'Yes',
                            handler: () => { doDeleteFile() }
                        },
                        {
                            title: 'No',
                            default: true,
                        },
                    ]
                });
            };
            return {
                state,
                propsHello,
                updateFileList,
                doDownload,
                doDeleteFile,
                doComfirmDelete
            };
        }
    });
</script>

<style scoped>
    * {
        box-sizing: border-box;
    }
    #list {
        max-width: 640px;
        margin: 0 auto;
    }
    .file-table {
        position: relative;
        top: 100px;
    }
    /*table {*/
    /*    width: 110%;*/
    /*    border-collapse: collapse;*/
    /*}*/
    /*thead th {*/
    /*    border-bottom: 2px solid #55aa42; !*#d31c4a *!*/
    /*    color: #429c43;*/
    /*}*/
    /*th,*/
    /*th {*/
    /*    padding: 0 8px;*/
    /*    line-height: 40px;*/
    /*}*/
    /*thead th.id {*/
    /*    width: 50px;*/
    /*}*/
    /*thead th.state {*/
    /*    width: 100px;*/
    /*}*/
    /*thead th.button {*/
    /*    width: 60px;*/
    /*}*/
    /*tbody td.button, tbody td.state {*/
    /*    text-align: center;*/
    /*}*/
    /*tbody tr td,*/
    /*tbody tr th {*/
    /*    border-bottom: 1px solid #ccc;*/
    /*    transition: all 0.4s;*/
    /*}*/
    /*tbody tr.done td,*/
    /*tbody tr.done th {*/
    /*    background: #f8f8f8;*/
    /*    color: #bbb;*/
    /*}*/
    /*tbody tr:hover td,*/
    /*tbody tr:hover th {*/
    /*    background: #f4fbff;*/
    /*}*/
    .table {
        text-align: center;
        font-weight: bold;
    }
    .download-button {
        border: none;
        border-radius: 20px;
        position: relative;
        top: 7px;
        line-height: 24px;
        padding: 0 8px;
        background: #0099e4;
        color: #fff;
        cursor: pointer;
    }
    .delete-button {
        border: none;
        border-radius: 20px;
        position: relative;
        top: 7px;
        left: 2px;
        line-height: 24px;
        padding: 0 8px;
        background: #e45856;
        color: #fff;
        cursor: pointer;
    }
    #post-file-loader {
        position: relative;
        top: 10px;
    }
</style>
