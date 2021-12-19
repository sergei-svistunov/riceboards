<template>
  <div class="d-inline-block">
    <span :title="hideName ? name || email : email">
      <img
          :src="avatarUrl"
          class="rounded-circle"
          :height="avatarSize"
          :alt="name"
          loading="lazy"
      />
    </span>
    <template v-if="!hideName">
      &nbsp;{{ name || email }}
    </template>
  </div>
</template>

<script lang="ts">
import {defineComponent} from 'vue';

export default defineComponent({
  name: 'BUser',
  props: {
    email: String,
    name: String,
    avatarUrl: String,
    avatarSize: String,
    hideName: Boolean
  },
});

</script>

