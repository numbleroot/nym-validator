<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
      <sp-type-form
        type="gateway"
        :fields="[
          'identityKey',
          'sphinxKey',
          'clientListener',
          'mixnetListener',
          'location',
        ]"
      />
      <sp-type-form
        type="mixnode"
        :fields="[
          'pubKey',
          'layer',
          'version',
          'host',
          'location',
          'reputation',
        ]"
      />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
