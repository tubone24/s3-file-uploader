<template>
  <div class="FileUpload" id="FileUpload">
    <div class="select">
      <b-field>
      <b-select id="format-select" v-model="state.selected">
        <option disabled value="">ファイル種類選択</option>
        <option value="test1">test1</option>
      </b-select>
        </b-field>
    </div>
      <div class="uploadButton">
        開く
        <input id="file-choice" type="file" @change="onFileChange" accept="text/csv">
      </div>
      <div v-if="state.isLoading">
        <div v-show="state.isLoading" id="post-file-loader">
          <vue-loading type="spiningDubbles" color="#333" :size="{ width: '200px', height: '200px' }"></vue-loading>
        </div>
      </div>
      <div v-if="state.selected && state.file">
        <b-button id="post-file" type="is-primary" @click="postFile">送信</b-button>
      </div>
  </div>
</template>

<script lang='ts'>
  import {
    createComponent,
    reactive,
    onBeforeMount,
    onMounted,
    computed,
    SetupContext,
    ref
  } from '@vue/composition-api';
  import axios from 'axios';
  import toast from '@nuxtjs/toast';
  import { VueLoading } from 'vue-loading-template';

  // data
  const state = reactive<{
    file: string;
    fileName: string,
    selected: string;
    isLoading: boolean;
  }>({
    file: '',
    fileName: '',
    selected: '',
    isLoading: false
  });

  const onFileChange = async (e: any): Promise<void> => {
    const files = e.target.files || e.dataTransfer.files;
    if (!files.length) {
      return;
    }
    state.fileName = e.target.files[0].name || e.dataTransfer.files[0].name;
    await createFile(files[0]);
  };

  const createFile = async (file: Blob): Promise<void> => {
    const reader = new FileReader();
    reader.onload = (e: any) => {
      state.file = e.target.result;
    };
    return new Promise((resolve, reject) => {
      reader.readAsDataURL(file);
      return resolve();
    });
  };

  export default createComponent({
    components: {
      VueLoading
    },
    setup (props, ctx:SetupContext) {
      const toast = ctx.root.$root.$toast;
      const postFile = async (e: any): Promise<void> => {

        state.isLoading = true;
        let statusCode;
        await axios.post('/upload', {
          fileType: state.selected,
          data: state.file,
          fileName: state.fileName,
        }).then((response) => {
          statusCode = response.status;
          toast.global.allSuccess();
        }).catch((err) => {
          statusCode = err.response.status;
          toast.global.unknownError();
        });
        state.isLoading = false;
      };


      return {
        state,
        onFileChange,
        postFile,
      }
    }
  })

</script>

<style scoped>
  .loader {
    color: #381eff;
    font-size: 90px;
    text-indent: -9999em;
    overflow: hidden;
    width: 1em;
    height: 1em;
    border-radius: 50%;
    margin: 72px auto;
    position: relative;
    -webkit-transform: translateZ(0);
    -ms-transform: translateZ(0);
    transform: translateZ(0);
    -webkit-animation: load6 1.7s infinite ease, round 1.7s infinite ease;
    animation: load6 1.7s infinite ease, round 1.7s infinite ease;
  }

  @-webkit-keyframes load6 {
    0% {
      box-shadow: 0 -0.83em 0 -0.4em, 0 -0.83em 0 -0.42em, 0 -0.83em 0 -0.44em, 0 -0.83em 0 -0.46em, 0 -0.83em 0 -0.477em;
    }
    5%,
    95% {
      box-shadow: 0 -0.83em 0 -0.4em, 0 -0.83em 0 -0.42em, 0 -0.83em 0 -0.44em, 0 -0.83em 0 -0.46em, 0 -0.83em 0 -0.477em;
    }
    10%,
    59% {
      box-shadow: 0 -0.83em 0 -0.4em, -0.087em -0.825em 0 -0.42em, -0.173em -0.812em 0 -0.44em, -0.256em -0.789em 0 -0.46em, -0.297em -0.775em 0 -0.477em;
    }
    20% {
      box-shadow: 0 -0.83em 0 -0.4em, -0.338em -0.758em 0 -0.42em, -0.555em -0.617em 0 -0.44em, -0.671em -0.488em 0 -0.46em, -0.749em -0.34em 0 -0.477em;
    }
    38% {
      box-shadow: 0 -0.83em 0 -0.4em, -0.377em -0.74em 0 -0.42em, -0.645em -0.522em 0 -0.44em, -0.775em -0.297em 0 -0.46em, -0.82em -0.09em 0 -0.477em;
    }
    100% {
      box-shadow: 0 -0.83em 0 -0.4em, 0 -0.83em 0 -0.42em, 0 -0.83em 0 -0.44em, 0 -0.83em 0 -0.46em, 0 -0.83em 0 -0.477em;
    }
  }

  @keyframes load6 {
    0% {
      box-shadow: 0 -0.83em 0 -0.4em, 0 -0.83em 0 -0.42em, 0 -0.83em 0 -0.44em, 0 -0.83em 0 -0.46em, 0 -0.83em 0 -0.477em;
    }
    5%,
    95% {
      box-shadow: 0 -0.83em 0 -0.4em, 0 -0.83em 0 -0.42em, 0 -0.83em 0 -0.44em, 0 -0.83em 0 -0.46em, 0 -0.83em 0 -0.477em;
    }
    10%,
    59% {
      box-shadow: 0 -0.83em 0 -0.4em, -0.087em -0.825em 0 -0.42em, -0.173em -0.812em 0 -0.44em, -0.256em -0.789em 0 -0.46em, -0.297em -0.775em 0 -0.477em;
    }
    20% {
      box-shadow: 0 -0.83em 0 -0.4em, -0.338em -0.758em 0 -0.42em, -0.555em -0.617em 0 -0.44em, -0.671em -0.488em 0 -0.46em, -0.749em -0.34em 0 -0.477em;
    }
    38% {
      box-shadow: 0 -0.83em 0 -0.4em, -0.377em -0.74em 0 -0.42em, -0.645em -0.522em 0 -0.44em, -0.775em -0.297em 0 -0.46em, -0.82em -0.09em 0 -0.477em;
    }
    100% {
      box-shadow: 0 -0.83em 0 -0.4em, 0 -0.83em 0 -0.42em, 0 -0.83em 0 -0.44em, 0 -0.83em 0 -0.46em, 0 -0.83em 0 -0.477em;
    }
  }

  @-webkit-keyframes round {
    0% {
      -webkit-transform: rotate(0deg);
      transform: rotate(0deg);
    }
    100% {
      -webkit-transform: rotate(360deg);
      transform: rotate(360deg);
    }
  }

  @keyframes round {
    0% {
      -webkit-transform: rotate(0deg);
      transform: rotate(0deg);
    }
    100% {
      -webkit-transform: rotate(360deg);
      transform: rotate(360deg);
    }
  }

  img {
    width: 30%;
    margin: auto;
    display: block;
    margin-bottom: 10px;
  }

  .select {
    margin: auto;
  }
  .uploadButton {
    display:inline-block;
    position:relative;
    overflow:hidden;
    border-radius:3px;
    background:#099;
    color:#fff;
    text-align:center;
    padding:10px;
    line-height:15px;
    width:70px;
    cursor:pointer;
  }
  .uploadButton:hover {
    background:#0aa;
  }
  .uploadButton input[type=file] {
    position:absolute;
    top:0;
    left:0;
    width:100%;
    height:100%;
    cursor:pointer;
    opacity:0;
  }
  .uploadValue {
    display:none;
    background:rgba(255,255,255,0.2);
    border-radius:3px;
    padding:3px;
    color:#ffffff;
  }
</style>
