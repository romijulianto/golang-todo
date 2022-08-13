<template>
    <el-row>
        <el-col style="width: 100%">
            <h1>Task Scheduler</h1>
            <todo-form @send-message="createTodo"></todo-form>
            <el-table :data="todos">
                <el-table-column prop="task" label="Task" width="350" />
                <el-table-column prop="employee" label="Employee" width="200" />
                <el-table-column prop="deadline" label="Deadline" width="200" />
                <el-table-column prop="description" label="Description" width="350" />
                <el-table-column prop="completed" label="Completed" width="200" />
                <el-table-column fixed="right" label="Operations" width="200">
                    <template #default="scope">
                        <el-space wrap>
                            <el-switch v-model="scope.row.completed" @click="updateTodo(scope.row)" />
                            <el-popconfirm confirm-button-text="Yes" cancel-button-text="No" icon="el-icon-info"
                                icon-color="red" title="Are you sure to delete the task ?"
                                @confirm="handleDelete(scope.row)" @cancel="cancelDetele">
                                <template #reference>
                                    <el-button size="mini" type="danger">Delete</el-button>
                                </template>
                                </el-popconfirm>
                                </el-space>
                    </template>
                </el-table-column>
            </el-table>
        </el-col>
    </el-row>
</template>

<script lang="ts">
import { ElMessage } from "element-plus";
import { Options, Vue } from "vue-class-component";
import TodoForm from "./TodoForm.vue";

interface Todo {
    id: number;
    task: string;
    employee: string;
    deadline: string;
    description: string;
    completed: boolean;
}

@Options({
    components: {
        TodoForm,
    }
})
export default class TodoList extends Vue {
    todos = [];

    async mounted() {
        await this.loadTodos();
    }

    async loadTodos() {
        const response = await this.axios.get(`http://localhost:3000/todos`);
        this.todos = response.data;
    }
    async createTodo(todo: Todo) {
        console.log("Task", todo);
        await this.axios.post(`http://localhost:3000/todos`, {
            task: todo.task,
            employee: todo.employee,
            deadline: todo.deadline,
            description: todo.description,
            completed: todo.completed
        });
        ElMessage({
            message: "Task Created",
            type: "success"
        });
        await this.loadTodos();
    }

    async updateTodo(todo: Todo) {
        console.log("Task", todo);
        await this.axios.post(`http://localhost:3000/todos/${todo.id}`, {
            id: todo.id,
            task: todo.task,
            employee: todo.employee,
            deadline: todo.deadline,
            description: todo.description,
            completed: todo.completed
        });
        ElMessage({
            message: "Task has been Updated",
            type: "success"
        });
        await this.loadTodos();
    }

    async handleDelete(todo: Todo) {
        await this.axios.delete(`http://localhost:3000/todos/${todo.id}`);
        ElMessage({
            message: "Task has been Deteled",
            type: "success"
        });
        await this.loadTodos();
    }

    cancelDetele() {
        console.log('Canceled the Delete')
    }
}
</script>