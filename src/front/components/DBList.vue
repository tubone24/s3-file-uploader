<template>
    <div id="dblist">
        <b-button id="search-file" type="is-primary" @click="updateDBList">更新</b-button>
        <div class="file-table">
            <b-table
               :data="state.dbList"
               :paginated="true"
               :pagination-position="both"
               :default-sort-direction="desc"
               default-sort="test"
            >
              <template slot-scope="dbList">
                <b-table-column field="id" label="id" sortable>
                    {{ dbList.row.id }}
                </b-table-column>
                <b-table-column field="test" label="test" sortable>
                  {{ dbList.row.test }}
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
    import {FileNotFoundError} from '~/types/error';
    import { dbKeyList, DbKeyList } from '@/config/config';
    // data
    const state = reactive<{
        dbList: Array<Map<string, string>>
        isLoading: boolean,
        dbKeyList: DbKeyList,
    }>({
        dbList: [],
        isLoading: false,
        dbKeyList: dbKeyList
    });


    const updateDBList = async (): Promise<void> => {
        const res = await axios.get('/api/checkdb');
        if (res.status === 200) {
            state.dbList = res.data;
        }
        //console.log(JSON.stringify(state.uploadList));
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
            onBeforeMount( async() => {
              await updateDBList()
            });
            return {
                state,
                propsHello,
                updateDBList
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
