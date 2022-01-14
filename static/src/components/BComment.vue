<template>
  <MDBTooltip v-model="show" style="cursor: help">
    <template #reference>
      <MDBIcon icon="comment" iconStyle="far"/>
    </template>
    <template #tip>
      <div >{{ text }}</div>
    </template>
  </MDBTooltip>
</template>

<script lang="ts">
import {defineComponent} from 'vue';
import {MDBIcon, MDBTooltip} from "mdb-vue-ui-kit";

export default defineComponent({
  name: 'BComment',
  components: {
    MDBIcon, MDBTooltip
  },
  props: {
    text: String,
  },
  data() {
    return {
      show: false
    }
  }
});
</script>

