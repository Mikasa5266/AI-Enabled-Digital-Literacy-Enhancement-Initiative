<template>
    <div>
        <div class="search">
            <InputSearch v-model:value="search_content" placeholder="请输入想要查询的内容" style="width: 200px;"
                @search="onsearchCards" />
        </div>
        <div style="padding: 20px;">
            <Row :gutter="16">
                <Col v-for="card in cards_present" :key="card.key" :span="4">
                    <Card :title="card.title" :bordered="false">
                        <template #cover>
                            <div class="cover_wrapper">
                                <div class="cover_mask">
                                    <Button type="primary" ghost @click="onclickpvButton">立即预览</Button>
                                </div>
                                <img class="cover_img" :alt="card.title" :src="card.cover_url" />
                            </div>
                        </template>
                    </Card>
                </Col>
            </Row>
        </div>
        <Modal v-model:visible="openPPTView" width="1200" :footer="null" centered title="课件预览"
            :bodyStyle="{ height: '70vh', padding: '0' }">
            <div class="w-3/5 h-full bg-gray-100">
                <iframe
                    :src="'https://view.officeapps.live.com/op/view.aspx?src=' + 'https://github.com/singpenguin/ppt/blob/master/CIKM-keynote-Nov2014.pdf'"
                    width="100%" height="100%" frameborder="0">
                    这是预览插件
                </iframe>
            </div>
        </Modal>
    </div>
</template>

<script setup lang="ts">
import { Button, Card, Col, InputSearch, Modal, Row } from 'ant-design-vue';
import { ref } from 'vue';
//数据结构定义
//搜索内容
const search_content = ref<string>('')
//pptcard接口
interface PPTCard {
    key: string
    cover_url: string
    title: string
}
const cards: PPTCard[] = [
    {
        key: '1',
        cover_url: new URL('../../images/test1.jpg', import.meta.url).href,
        title: '如何打王者荣耀'
    },
    {
        key: '2',
        cover_url: new URL('../../images/test2.jpg', import.meta.url).href,
        title: '王者荣耀英雄选择技巧'
    },
    {
        key: '3',
        cover_url: new URL('../../images/test3.jpg', import.meta.url).href,
        title: '王者荣耀铭文搭配指南'
    },
    {
        key: '4',
        cover_url: new URL('../../images/test4.jpg', import.meta.url).href,
        title: '王者荣耀打野节奏把控'
    },
    {
        key: '5',
        cover_url: new URL('../../images/test5.jpg', import.meta.url).href,
        title: '王者荣耀中路支援技巧'
    },
    {
        key: '6',
        cover_url: new URL('../../images/test6.jpg', import.meta.url).href,
        title: '王者荣耀下路双人组配合'
    },
    {
        key: '7',
        cover_url: new URL('../../images/test7.jpg', import.meta.url).href,
        title: '王者荣耀上单抗压思路'
    },
    {
        key: '8',
        cover_url: new URL('../../images/test8.jpg', import.meta.url).href,
        title: '王者荣耀团战站位技巧'
    },
    {
        key: '9',
        cover_url: new URL('../../images/test9.jpg', import.meta.url).href,
        title: '王者荣耀逆风局翻盘攻略'
    },
    {
        key: '10',
        cover_url: new URL('../../images/test10.jpg', import.meta.url).href,
        title: '王者荣耀常用道具使用时机'
    }
]
const cards_present = ref<PPTCard[]>(cards)

//打开ppt模态框
const openPPTView = ref<boolean>(false)

//处理预览按钮
const onclickpvButton = () => {
    console.log('我是预览')
    openPPTView.value = true
    console.log(openPPTView.value)
}
//处理搜索
const onsearchCards = () => {
    const search_result: PPTCard[] = cards.filter(card => {
        return card.title.includes(search_content.value)
    })
    console.log(search_content.value)
    console.log(search_result)
    cards_present.value = search_result
}


</script>

<style scoped>
.cover_wrapper {
    position: relative;
    width: 100%;
    aspect-ratio: 16/9;
    overflow: hidden;
    cursor: pointer;
}

.cover_img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: transform 0.3s ease;
}

.cover_mask {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    opacity: 0;
    transition: opacity 0.3s ease;
}

.cover_wrapper:hover .cover_mask {
    opacity: 1;
}
</style>