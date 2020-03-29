<template>
    <div id="list">
        <div class="select">
            <b-field>
                <b-select id="format-select" v-model="state.selected">
                    <option disabled value="">ファイル種類選択</option>
                    <option value="test-test">test-test</option>
                </b-select>
            </b-field>
            <div v-if="state.selected">
                <b-button id="search-file" type="is-primary" @click="updateFileList">検索</b-button>
            </div>
        </div>
        <div class="file-table">
        <table>
            <thead>
            <tr>
                <th class="filepath">ファイル名</th>
                <th class="updated">更新日時</th>
                <th class="size">サイズ</th>
                <th class="download">-</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="item in state.uploadList">
                <th>{{ item.name }}</th>
                <th>{{ item.lastModified }}</th>
                <th>{{ item.size }}</th>
                <button v-on:click="doDownload(item.name)">DL</button>
            </tr>
            </tbody>
        </table>
        </div>
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
    import {FileNotFoundError} from "~/types/error";
    // data
    const state = reactive<{
        uploadList: Array<Map<string, string>>
        selected: string
    }>({
        uploadList: [],
        selected: ''
    });

    const updateFileList = async (): Promise<void> => {
        const res = await axios.get('/api/list?prefix=' + state.selected);
        if (res.status === 200) {
            state.uploadList = res.data.fileList;
        }
        console.log(JSON.stringify(state.uploadList));
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
        props: {
            propHello: {
                type: String
            }
        },
        setup (props: Props, ctx:SetupContext) {
            // props
            const propsHello = props.propHello;
            const toast = ctx.root.$root.$toast;
            const doDownload = async (filePath: string): Promise<void> => {
                try{
                    const blob = await downloadFile(filePath);
                    const link = document.createElement('a');
                    link.href = window.URL.createObjectURL(blob);
                    const filename = filePath.match(".+/(.+?)([\?#;].*)?$")[1];
                    link.download = filename;
                    link.click();
                } catch (e) {
                    if (e instanceof FileNotFoundError) {
                        console.log(toast);
                        toast.global.nofileError();
                    } else {
                        toast.global.unknownError();
                    }
                }
            };
            return {
                state,
                propsHello,
                updateFileList,
                doDownload
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
    table {
        width: 110%;
        border-collapse: collapse;
    }
    thead th {
        border-bottom: 2px solid #55aa42; /*#d31c4a */
        color: #429c43;
    }
    th,
    th {
        padding: 0 8px;
        line-height: 40px;
    }
    thead th.id {
        width: 50px;
    }
    thead th.state {
        width: 100px;
    }
    thead th.button {
        width: 60px;
    }
    tbody td.button, tbody td.state {
        text-align: center;
    }
    tbody tr td,
    tbody tr th {
        border-bottom: 1px solid #ccc;
        transition: all 0.4s;
    }
    tbody tr.done td,
    tbody tr.done th {
        background: #f8f8f8;
        color: #bbb;
    }
    tbody tr:hover td,
    tbody tr:hover th {
        background: #f4fbff;
    }
    button {
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
</style>
