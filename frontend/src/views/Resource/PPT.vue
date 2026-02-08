<template>
    <div>
        <div class="search">
            <InputSearch v-model:value="search_content" placeholder="请输入想要查询的内容" style="width: 200px;"
                @search="onsearchCards" />
        </div>
        <div style="padding: 20px;">
            <Row :gutter="[16,24]">
                <Col v-for="card in cards_present" :key="card.id" :span="4">
                    <Card :title="card.title" :bordered="false">
                        <template #cover>
                            <div class="cover_wrapper">
                                <div class="cover_mask">
                                    <Button type="primary" ghost @click="onclickpvButton(card)">立即预览</Button>
                                </div>
                                <img class="cover_img" :alt="card.title" :src="card.cover_url" />
                            </div>
                        </template>
                    </Card>
                </Col>
            </Row>
        </div>
        <Modal v-model:open="openPPTView" :width="1200" :footer="null" 
            :bodyStyle="{  height: '75vh', padding: '0', overflow: 'hidden' }">
            <div style="width: 100%; height: 100%; background-color: #f3f4f6;" class="w-full h-full bg-gray-100">
                <iframe :src="`/pdfjs/web/viewer.html?file=${encodeURIComponent(currentPdfUrl)}#view=FitH`" width="100%"
                    height="100%" frameborder="0">
                    这是预览插件
                </iframe>
            </div>
        </Modal>
    </div>
</template>

<script setup lang="ts">
import { Button, Card, Col, InputSearch, Modal, Row } from 'ant-design-vue';
import { onMounted, ref } from 'vue';
import type { PPTcard } from '../../types/ppt';
import axios from 'axios';

// 后端返回的包装结构
interface ApiResponse<T> {
    code: number
    msg: string
    data: T
}

//搜索内容
const search_content = ref<string>('')
//pptcard数组
let cards: PPTcard[] = []
const cards_present = ref<PPTcard[]>([])

//获取内容
const getPPT = async () => {
    try {
        const result = await axios.get<ApiResponse<PPTcard[]>>("/api/ppt")
        cards = result.data.data
        cards_present.value = cards  // 更新响应式数据，页面才会渲染
        console.log(cards)
        console.log(cards_present.value)
    } catch (error) {
        console.log("请求错误", error)
    }
}
onMounted(()=>{
    getPPT()
})
//打开ppt模态框
const openPPTView = ref<boolean>(false)
//当前pdf的URL
const currentPdfUrl = ref<string>('')
//处理预览按钮
const onclickpvButton = (card: PPTcard) => {
    console.log('我是预览')
    currentPdfUrl.value = card.file_url
    openPPTView.value = true
    console.log(openPPTView.value)
}
//处理搜索
const onsearchCards = () => {
    const search_result: PPTcard[] = cards.filter(card => {
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
<style scoped>
:deep(.ant-card){
    border-radius: 12px;
    overflow: hidden;
    box-shadow: 0 2px 8px rgba(0,0,0,0.1);
    transition:  box-shadow 0.3s ease;
}
:deep(.ant-card:hover){
    box-shadow:0 4px 16px rgba(0,0,0,0.2);
}
</style>