<template>
    <div style="height: 100%">
        <div class="header">
            <div class="search">
                <InputGroup compact>
                    <Select v-model:value="search_content.category" style="width: 40%;" size="large">
                        <Option class="selct_option" value="全部">全部</Option>
                        <Option class="select_option" value="前端开发">前端开发</Option>
                        <Option class="select_option" value="移动端开发">移动端开发</Option>
                        <Option class="select_option" value="后端开发">后端开发</Option>
                    </Select>
                    <InputSearch v-model:value="search_content.title" placeholder="请输入想要查询的内容"
                        style="max-width: 450px;width: 60%;" enterButton size="large" @search="onsearchCards" />
                </InputGroup>
            </div>
            <div class="upload">
                <Button @click="handleUpload" type="primary" size="large">
                    <UploadOutlined />
                    上传PPT
                </Button>
            </div>
        </div>
        <div class="main">
            <div v-if="cards_present.length" style="padding: 20px;">
                <Row :gutter="[16, 24]">
                    <Col v-for="card in cards_present" :key="card.id" :xs="24" :sm="12" :md="8" :xl="6">
                        <Card :bordered="false">
                            <template #cover>
                                <div class="cover_wrapper">
                                    <div class="cover_mask">
                                        <Button type="primary" ghost @click="onclickpvButton(card)">立即预览</Button>
                                    </div>
                                    <img class="cover_img" :alt="card.title" :src="card.cover_url" />
                                </div>
                                <div class="title_tag">
                                    <div class="title">
                                        <span>{{ card.title }}</span>
                                    </div>
                                    <div>
                                        <Tag class="tags" :bordered="false" :color="TagColorMap[card.category]">
                                            {{ card.category }}
                                        </Tag>
                                    </div>
                                </div>
                            </template>
                        </Card>
                    </Col>
                </Row>
            </div>
            <div v-else style="display: flex; justify-content: center; align-items: center; flex-direction: column;">
                <Empty />
                <div>赶快上传你的ppt吧!</div>
            </div>
        </div>
        <Modal v-model:open="openPPTView" :width="1200" :footer="null"
            :bodyStyle="{ height: '75vh', padding: '0', overflow: 'hidden' }">
            <div style="width: 100%; height: 100%; background-color: #f3f4f6;" class="w-full h-full bg-gray-100">
                <iframe :src="currentPdfUrl" width="100%" height="100%" frameborder="0">
                    这是预览插件
                </iframe>
            </div>
        </Modal>
        <Modal v-model:open="openUploadView" :width="600" :footer="null"
            :bodyStyle="{ height: '65vh', padding: '0', overflow: 'hidden' }">
            <div style="display: flex; align-items: center; justify-content:center;padding-top: 20px;">
                <Form ref="uploadFormRef" :model="uploadPPT" layout="vertical" labelAlign="left" :requiredMark="false" @finish="onFinish">
                    <FormItem name="title" label="输入标题" style="margin-bottom: 20px;" :rules="[{ required: true, message: '标题为必填' }]">
                        <Input v-model:value="uploadPPT.title" size="large" style="border-radius: 8px;"
                            :maxlength="50" />
                    </FormItem>
                    <FormItem name="category" label=" 选择分类" style="margin-bottom: 20px;"
                        :rules="[{ required: true, message: '分类为必填' }]">
                        <Select v-model:value="uploadPPT.category" size="large" style="border-radius: 8px;">
                            <Option class="select_option" value="全部">全部</Option>
                            <Option class="select_option" value="前端开发">前端开发</Option>
                            <Option class="select_option" value="移动端开发">移动端开发</Option>
                            <Option class="select_option" value="后端开发">后端开发</Option>
                        </Select>
                    </FormItem>
                    <FormItem name="coverImgList" label="点击上传封面图(png,<=1MB)" style="margin-bottom: 20px;"
                        :rules="[{ required: true, message: '必须选择封面图' }]">
                        <Upload v-model:fileList="uploadPPT.coverImgList" listType="picture-card" :maxCount="1"
                            :beforeUpload="beforecover">
                            <Image v-if="uploadPPT.coverImgList[0]?.url" :src="uploadPPT.coverImgList[0]?.url" alt="封面">
                            </Image>
                            <div>
                                <loading-outlined v-if="loading"></loading-outlined>
                                <plus-outlined v-else></plus-outlined>
                                <div class="ant-upload-text">Upload</div>
                            </div>
                        </Upload>
                    </FormItem>
                    <FormItem name="filelist" label="选择文件(pdf)" style="margin-bottom: 20px;"
                        :rules="[{ required: true, message: '必须选择ppt' }]">
                        <Upload action="javascript:;" v-model:fileList="uploadPPT.filelist" :beforeUpload="beforeppt"
                            :maxCount="1" :auto-upload="false">
                            <Button type="primary" size="large" style="width: 300px;">
                                上传PPT
                            </Button>
                        </Upload>
                    </FormItem>
                    <FormItem>
                        <Button type="primary" size="large" style="width: 300px;" htmlType="submit">确认</Button>
                    </FormItem>
                    <FormItem>
                        <Button type="primary" size="large" style="width: 300px;" @click="cancelSubmit">退出</Button>
                    </FormItem>
                </Form>
            </div>
        </Modal>
        <Modal v-model:open="openConfirmCancel" centered width="400" ok-text="确认退出" cancel-text="继续填写"
            @ok="onConfirmExit">
            确认要取消上传吗?<br />未提交的数据将会消失
        </Modal>
    </div>
</template>

<script setup lang="ts">
import { Button, Card, Col, Empty, Form, FormItem, Image, Input, InputGroup, InputSearch, message, Modal, Row, Select, Tag, Upload, type UploadProps } from 'ant-design-vue';
import { onMounted, ref } from 'vue';
import type { ApiResponse, Category, PPTcard, SearchContent, UploadPPT } from '../../types/ppt';
import axios from 'axios';
import { Option } from 'ant-design-vue/es/vc-select';
import { transformURL } from '../../utils/transformURL';
import { UploadOutlined } from '@ant-design/icons-vue';
// ---------------数据定义---------//
//搜索内容
const search_content = ref<SearchContent>({
    category: '全部',
    title: ''
})
//pptcard数组
let cards: PPTcard[] = []
const cards_present = ref<PPTcard[]>([])
//不同类别标签所用的颜色表
const TagColorMap: Record<Category, string> = {
    '全部': 'white',
    '前端开发': 'red',
    '后端开发': 'blue',
    '移动端开发': 'yellow',
    '': 'white'
}
//上传的PPT的表单元素
const uploadPPT = ref<UploadPPT>({
    title: '',
    category: '',
    coverImgList: [],
    filelist: []
})

//定义 Form 的 ref 引用
const uploadFormRef = ref<typeof Form | null>(null)
//最大上传大小
const MAX_FILESIZE = 1024 * 1024
//打开ppt模态框
const openPPTView = ref<boolean>(false)
//当前pdf的URL
const currentPdfUrl = ref<string>('')
//打开Upload模态框
const openUploadView = ref<boolean>(false)
//图片加载中
const loading = ref<boolean>(false);
//确认是否取消的模态框
const openConfirmCancel = ref<boolean>(false)
// ---------------事件处理----------------//
//处理预览按钮
const onclickpvButton = (card: PPTcard) => {
    console.log('我是预览')
    currentPdfUrl.value = card.file_url
    openPPTView.value = true
    console.log(openPPTView.value)
}
//处理上传页面打开按钮
const handleUpload = () => {
    openUploadView.value = true
}
//处理搜索
const onsearchCards = () => {
    const search_result: PPTcard[] = cards.filter(card => {
        return card.title.includes(search_content.value.title) && (card.category.includes(search_content.value.category) || search_content.value.category === '全部')
    })
    console.log(search_content.value)
    console.log(search_result)
    cards_present.value = search_result
}
//上传前的文件处理（cover）
const beforecover: UploadProps['beforeUpload'] = file => {
    const ispng = file.type === 'image/png'
    if (!ispng) {
        message.error(`${file.name}不是png文件`)
        return Upload.LIST_IGNORE
    }
    const isoverbig = file.size > MAX_FILESIZE
    if (isoverbig) {
        message.error(`文件${file.name}过大`)
        return Upload.LIST_IGNORE
    }
    return false
}
//上传前的文件处理（ppt）
const beforeppt: UploadProps['beforeUpload'] = file => {
    const ispdf = file.type === 'application/pdf'
    if (!ispdf) {
        message.error(`${file.name}不是pdf文件`)
        return Upload.LIST_IGNORE
    }
    return false
}
const onFinish = () => {
    const formData = new FormData()
    formData.append('title', uploadPPT.value.title)
    formData.append('category', uploadPPT.value.category)

    // --- 修复封面图上传 ---
    if (uploadPPT.value.coverImgList.length > 0) {
        // 获取 AntD 封装对象中的原始文件对象 originFileObj
        const coverItem = uploadPPT.value.coverImgList[0]
        if (coverItem && coverItem.originFileObj) {
            formData.append('coverFile', coverItem.originFileObj)
        }
    }

    // --- 修复 PPT 文件上传 ---
    if (uploadPPT.value.filelist.length > 0) {
        // 获取 AntD 封装对象中的原始文件对象 originFileObj
        const pptItem = uploadPPT.value.filelist[0]
        if (pptItem && pptItem.originFileObj) {
            formData.append('pptFile', pptItem.originFileObj)
        }
    }

    createPPT(formData)
}
const cancelSubmit = () => {
    openConfirmCancel.value = true
}
// --------------------接口请求----------------
//获取内容
const getPPT = async () => {
    try {
        const result = await axios.get<ApiResponse<PPTcard[]>>("/api/ppt")
        console.log(result)
        cards = result.data.data
        console.log(cards)
        cards_present.value = cards  // 更新响应式数据，页面才会渲染
        cards_present.value.forEach(card => {
            card.file_url = transformURL(card.file_url)
            card.cover_url = transformURL(card.cover_url)
        });
        console.log(cards)
        console.log(cards_present.value)
    } catch (error) {
        console.log("请求错误", error)
    }
}
const createPPT = async (formData: FormData) => {
    try {
        const result = await axios.post("/api/ppt", formData)
        if (result.data.code === 200) {
            message.success('PPT上传成功')
            uploadPPT.value = {
                title: '',
                category: '',
                coverImgList: [],
                filelist: []
            }
            openUploadView.value = false
            await getPPT()
        } else {
            message.error(`上传失败：${result.data.msg || '未知错误'}`)
        }
    } catch (error: any) {
        // 关键修改：打印完整的 error 对象
        console.error("上传报错详情:", error)

        if (error.response) {
            // 请求已发出，服务器返回了状态码（非 2xx）
            console.log("状态码:", error.response.status);
            console.log("后端返回数据:", error.response.data);
            message.error(`上传失败: ${error.response.status} ${error.response.data?.message || ''}`);
        } else {
            message.error('网络错误或请求未发送');
        }
    }
}
const onConfirmExit = () => {
    if (uploadFormRef.value) {
        uploadFormRef.value.resetFields()
    }
    openUploadView.value = false
    openConfirmCancel.value = false
}
// --------------------生命周期-------------------//
onMounted(() => {
    getPPT()
})
</script>
<!-- 头部的样式 -->
<style scoped>
.header {
    display: flex;
    justify-content: center;
    align-items: center;
    max-height: 300px;
    height: 10%;
    padding: 20px;
    background-image: linear-gradient(to right, #2645d0, #12a2a0);
}

.search {
    max-width: 500px;
    width: 70%;
}

.upload {
    margin-left: 20px;
    display: flex;
    align-items: center;
}
</style>
<!-- 卡片区域的样式 -->
<style scoped>
.main {
    height: 90%;
    background-color: #f3f4f6;
}

.cover_wrapper {
    position: relative;
    width: 100%;
    aspect-ratio: 16/9;
    overflow: hidden;
    cursor: pointer;
    border-radius: 10px;
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
    background-color: rgba(0, 0, 0, 0.45);
}

.cover_wrapper:hover .cover_mask {
    opacity: 1;
}
.title_tag{
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.title {
    font-size: 28px;
    padding: 12px;
}

.tags {
    font-size: 16px;
    padding: 10px;
}
</style>
<style scoped>
:deep(.ant-card) {
    border-radius: 12px;
    overflow: hidden;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    transition: box-shadow 0.3s ease;
}

:deep(.ant-card:hover) {
    box-shadow: 0 4px 16px rgba(0, 0, 0, 0.2);
}
</style>