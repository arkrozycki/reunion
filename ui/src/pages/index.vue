<script setup lang="ts">
import axios, { AxiosResponse, AxiosError } from 'axios'
import { reactive } from 'vue'

const locationURI =
  location.host == 'localhost:8088' ? 'localhost:8089' : location.host
const protocol = 'http://'
const port = 9011

const userObj = reactive({ email: '', password: '' })

function RegisterUser() {
        if (userObj.email == '' || userObj.password == ''){
                console.log('fill out form')
                return 
        }

        console.log(userObj.email, userObj.password)
}
</script>

<template>
  <Base>
  <div class="min-h-full flex flex-col justify-center py-12 sm:px-6 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-md">
      <h2 class="mt-6 -mb-5 -ml-12 text-center text-xl text-gray-900">
        Sign up for
      </h2>
      <h1 class="text-indigo-600 text-center text-6xl font-extrabold">Reunion</h1>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
      <div class="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
        <form
          class="space-y-6"
          action="#"
          method="POST"
        >
          <div>
            <label
              for="email"
              class="block text-sm font-medium text-gray-700"
            >
              Email address
            </label>
            <div class="mt-1">
              <input
                v-model="userObj.email"
                id="email"
                name="email"
                type="email"
                autocomplete="email"
                required=""
                class="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              />
            </div>
          </div>

          <div>
            <label
              for="password"
              class="block text-sm font-medium text-gray-700"
            >
              Password
            </label>
            <div class="mt-1">
              <input
                v-model="userObj.password"
                id="password"
                name="password"
                type="password"
                autocomplete="current-password"
                required=""
                class="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
              />
            </div>
          </div>

          <div>
            <button
              @click="RegisterUser"
              type="button"
              class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Register
            </button>
          </div>
        </form>

      </div>
    </div>
  </div>
  </Base>
</template>

<style></style>
