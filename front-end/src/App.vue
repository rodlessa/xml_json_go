<template>
  <div id="app" class="flex flex-col items-center justify-center h-screen p-4 font-sans bg-white">
    <h1 class="text-2xl font-bold text-center mb-4 text-gray-800">XML para JSON Converter</h1>

    <div class="flex space-x-4 w-full mb-4">
      <div class="flex-1">
        <label for="xmlInput" class="block text-gray-700">XML Input:</label>
        <textarea
          id="xmlInput"
          v-model="xmlInput"
          @input="convertXmlToJson"
          placeholder="Insira o XML aqui..."
          rows="10"
          class="w-full h-80 p-2 border border-gray-300 rounded-md text-gray-900"
        ></textarea>
      </div>

      <div class="flex-1">
        <label for="jsonOutput" class="block text-gray-700">JSON Output:</label>
        <textarea
          id="jsonOutput"
          v-model="jsonOutput"
          readonly
          rows="10"
          class="w-full h-80 p-2 border border-gray-300 rounded-md text-gray-900"
        ></textarea>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      xmlInput: '',
      jsonOutput: '',
    };
  },
  methods: {
    async convertXmlToJson() {
      try {
        const response = await axios.post('https://xmljson.rdls.dev/convert', this.xmlInput, {
          headers: { 'Content-Type': 'application/xml' },
        });
        this.jsonOutput = JSON.stringify(response.data, null, 2);
      } catch (error) {
        console.error('Erro ao converter XML para JSON:', error);
        this.jsonOutput = 'Erro na conversão.';
      }
    },
  },
};
</script>

<style>
</style>
