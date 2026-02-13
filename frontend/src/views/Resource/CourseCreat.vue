<template>
    <div class="wrapper_all">
        <header>
            <div>
                <h1>AI课程创作助手</h1>
                <br />
                <p>基于混元大模型的智能课程生成</p>
            </div>
        </header>
        <main>
            <div class="left">
                <Form :model="course" name="course" @finish="onFinish" @finish-failed="onFinishFail">
                    <h2>课程需求</h2>
                    <FormItem name="theme" :rules="[{ required: true, message: '课程主题为必填' }]">
                        <label class="form_label">
                            课程主题
                        </label>
                        <Input style="width: 350px;" v-model:value="course.theme" placeholder="例如:python编程入门" />
                    </FormItem>
                    <FormItem name="targetmember" :rules="[{ required: true, message: '目标受众为必选' }]">
                        <label class="form_label">
                            目标受众
                        </label>
                        <Select style="width: 350px;" v-model:value="course.targetmember">
                            <Option value="grade_one">一年级</Option>
                            <Option value="grade_two">二年级</Option>
                            <Option value="grade_three">三年级</Option>
                        </Select>
                    </FormItem>
                    <FormItem name="timelenth">
                        <label>
                            课程时长(分钟)
                        </label>
                        <Input style="width: 350px;" v-model:value="course.timelenth" />
                    </FormItem>
                    <FormItem name="chapterlenth">
                        <label class="form_label">
                            建议章节数
                        </label>
                        <Slider range v-model:value="course.chapterlenth" :max="10" />
                    </FormItem>
                    <FormItem>
                        <label class="form_label">
                            参考ppt资源(可选)
                        </label>
                        <Select style="width: 350px;" />
                    </FormItem>
                    <FormItem name="introduction">
                        <label class="form_label">
                            补充说明
                        </label>
                        <Textarea v-model:value="course.introduction" style="width: 350px;" :rows="4" placeholder="其他特殊要求或期望的内容方向" />
                    </FormItem>
                    <FormItem>
                        <Button type="primary" html-type="submit" style="width: 350px; height: 40px;    background-image: linear-gradient(to bottom right, #667de9, #754da5);">
                            生成课程大纲
                        </Button>
                    </FormItem>
                </Form>
            </div>
            <div class="right">
                <div style="width: 700px;">填写课程信息后,点击生成按钮开始创作</div>
            </div>
        </main>
    </div>
</template>

<script setup lang="ts">
import { Button, Form, FormItem, Input, Select, Slider, Textarea } from 'ant-design-vue';
import { ref } from 'vue';
import type { Course_Create_Info } from '../../types/course_create';
import { Option } from 'ant-design-vue/es/vc-select';
//表单元素
const course = ref<Course_Create_Info>({
    theme: '',
    targetmember: '',
    timelenth: 0,
    chapterlenth: [3, 4],
    inferresourse: '',
    introduction: ''
})


//提交表单
const onFinish = () => {

}
const onFinishFail = () => {

}
</script>

<style scoped>
.wrapper_all {
    display: flex;
    flex-direction: column;
    width: 100%;
    height: 100%;
    background-image: linear-gradient(to bottom right, #667de9, #754da5);
}

header {
    color: white;
    height: 150px;
    display: flex;
    align-items: center;
    justify-content: center;
}

main {
    flex: 1;
    display: flex;
    justify-content: space-evenly;
}

.left {
    height: 700px;
    width: 400px;
    border-radius: 20px;
    background-color: white;
    display: flex;
    flex-direction: column;
    gap: 20px;
}

.right {
    border-radius: 20px;
    background-color: white;
    height: 700px;
    width: 800px;
    display: flex;
    align-items: center;
    justify-self: center;
}

.form_label {
    display: block;
}

.left h2 {
    margin-top: 10px;
    margin-bottom: 10px;
}
</style>