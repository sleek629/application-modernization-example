<template>
  <div id='app'>
    <section>
      <span class='title-text'> Word Database</span>
      <div class='row justify-content-center mt-4'>
        <input v-model='inputField' v-on:keyup.enter='updateWord' class='mr-1' placeholder='Input word'>
        <button @click='updateWord()' class='btn btn-primary'>Submit</button>
      </div>
    </section>
    <section>
      <template v-if="inputField">
        Your input: {{ inputField }}
      </template>
    </section>
    <section>
      <b-table striped hover :items="wordCounts" :fileds="fields">
      </b-table>
    </section>
  </div>
</template>

<script>
import { Empty, Word } from './word_pb'
import { WordAPIClient } from './word_grpc_web_pb'
export default {
  name: 'app',
  components: {},
  data: function () {
    return {
      inputField: '',
      wordCounts: [],
      fields:[
      {key: 'word', label: 'Word'},
      {key: 'count', label: 'Count'},
      ]
    }
  },
  created: function () {
    // eslint-disable-next-line
    // this.client = new WordAPIClient("http://localhost:8080", null, null)
    this.client = new WordAPIClient(process.env.VUE_APP_WORDAPI_ADDRESS, null, null)
    this.getWords()
  },
  methods: {
    getWords: function () {
      // eslint-disable-next-line
      let getRequest = new Empty()
      // eslint-disable-next-line
      this.client.getWords(getRequest, {}, (err, response) => {
        this.wordCounts = response.toObject().wcList
      })
    },
    updateWord: function () {
      // eslint-disable-next-line
      let postRequest = new Word()
      postRequest.setWord(this.inputField)
      // eslint-disable-next-line
      this.client.updateWord(postRequest, {}, (err, response) => {
          this.getWords()
      })
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.title-text {
  font-size: 22px;
}
</style>
