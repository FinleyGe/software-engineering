import {inject} from "vue";
import type {VueCookies} from "vue-cookies";

export function useVueCookies(): VueCookies {
  const $cookies = inject<VueCookies>("$cookies");
  return $cookies as VueCookies;
}
