import type {VueCookies} from "vue-cookies";
import {inject} from "vue";

export function useVueCookies(): VueCookies {
  const $cookies = inject<VueCookies>("$cookies");
  return $cookies as VueCookies;
}
