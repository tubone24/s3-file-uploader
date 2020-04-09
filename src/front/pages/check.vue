<template>
    <section class="section">
        <div class="container">
            <db-list  propHello="HelloComposition"/>
        </div>
    </section>
</template>

<script lang="ts">
    import {
        createComponent,
        reactive,
        onMounted,
        computed,
        ref
    } from '@vue/composition-api';
    import DbList from '@/components/DBList.vue';
    export default createComponent({
        layout: 'default',
        components: {
          DbList
        },
        setup() {
        }
    })
</script>
