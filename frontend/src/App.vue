<template>
  <div class="container">
    <h1 class="display-1">Reminder App</h1>
    <form @submit.prevent="addTodo">
      <div class="input-group mb-3">
        <input type="text"
               class="form-control"
               placeholder="Enter a new reminder"
               aria-label="Enter a new reminder"
               aria-describedby="button-add"
               v-model="newTodoDescription">
        <button class="btn btn-outline-primary" type="submit" id="button-add">Add Reminder</button>
      </div>
    </form>
    <ul class="list-group mt-3">
      <li class="d-flex justify-content-between list-group-item list-group-item-action align-items-center"
          v-for="todo in todos"
          :key="todo.id">
          <span>
            <input class="form-check-input me-1"
                   type="checkbox"
                   value=""
                   aria-label="..."
                   v-model="todo.done"
                   @change="updateTodoStatus(todo)">
            <span :class="{ 'completed': todo.done }">{{ todo.task }}</span>
          </span>
        <span>
            <button class="btn" @click="deleteTodo(todo.id)">🗑</button>
          </span>
      </li>
    </ul>
  </div>
</template>

<script>
import {onMounted, ref} from 'vue';
import axios from 'axios';

export default {
  setup() {
    const todos = ref([]);
    const newTodoDescription = ref('');

    const fetchTodos = async () => {
      try {
        const response = await axios.get('http://localhost:8080/api/reminder/all');
        todos.value = response.data.map(todo => ({
          id: todo.id,
          task: todo.task,
          done: todo.done
        }));
      } catch (error) {
        console.error(error);
      }
    };

    const addTodo = async () => {
      if (newTodoDescription.value.trim() !== '') {
        const newTodo = {
          task: newTodoDescription.value.trim(),
          done: false
        };
        try {
          const response = await axios.post('http://localhost:8080/api/reminder/', newTodo);
          const createdTodo = response.data;
          todos.value.push(createdTodo);
          newTodoDescription.value = '';
        } catch (error) {
          console.error(error);
        }
      }
    };

    const deleteTodo = async (id) => {
      try {
        await axios.delete(`http://localhost:8080/api/reminder/${id}`);
        todos.value = todos.value.filter(todo => todo.id !== id);
      } catch (error) {
        console.error(error);
      }
    };

    const updateTodoStatus = async (todo) => {
      try {
        await axios.put('http://localhost:8080/api/reminder/', {id: todo.id});
      } catch (error) {
        console.error(error);
      }
    };

    onMounted(fetchTodos);

    return {
      todos,
      newTodoDescription,
      addTodo,
      deleteTodo,
      updateTodoStatus
    };
  }
};
</script>

<style>
.completed {
  text-decoration: line-through;
}
</style>
