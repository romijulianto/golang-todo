<template>
    <el-form @submit.prevent :model="formInput" class="task-form-input" :label-position="labelPosition">
        <el-form-item label="Task Title" prop="task">
            <el-input v-model="formInput.task" placeholder="Input Task"></el-input>
        </el-form-item>
        <el-form-item label="Employee" prop="employee">
            <el-input v-model="formInput.employee" placeholder="Input Employee"></el-input>
        </el-form-item>
        <el-form-item label="Deadline" prop="deadline">
            <el-input v-model="formInput.deadline" placeholder="MM/DD/YYY"></el-input>
        </el-form-item>
        <el-form-item label="Description" prop="description">
            <el-input v-model="formInput.description" placeholder="Description"/>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="onSubmit()">Add Task</el-button>
        </el-form-item>
    </el-form>
</template>

<script lang="ts">
import { ElMessage } from "element-plus";
import { Options, Vue } from "vue-class-component";

@Options({})
export default class TodoForm extends Vue {

    labelPosition = ('top')

    formInput = {
        task: "",
        employee: "",
        deadline: "",
        description: "",
        completed: false,
    };

    onSubmit() {
        if (this.formInput.task.length > 3) {
            this.$emit('send-message', this.formInput);
        } else {
            ElMessage({
                message: "Warning, This Task is to short",
                type: "warning"
            })
            this.formInput.task = "";
            this.formInput.employee = "";
            this.formInput.deadline = "";
            this.formInput.description = "";
        }
    }
}
</script>