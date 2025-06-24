const opt = Object.prototype.toString;
function isObject(obj) {
  return opt.call(obj) === "[object Object]";
}
function isString(obj) {
  return opt.call(obj) === "[object String]";
}
function isFunction(obj) {
  return typeof obj === "function";
}
function goodTimeText() {
  const time = /* @__PURE__ */ new Date();
  const hour = time.getHours();
  return hour < 9 ? "早上好" : hour <= 11 ? "上午好" : hour <= 13 ? "中午好" : hour <= 18 ? "下午好" : "晚上好";
}
function setObjToUrlParams(baseUrl, obj) {
  let parameters = "";
  for (const key in obj) {
    parameters += key + "=" + encodeURIComponent(obj[key]) + "&";
  }
  parameters = parameters.replace(/&$/, "");
  return /\?$/.test(baseUrl) ? baseUrl + parameters : baseUrl.replace(/\/?$/, "?") + parameters;
}
function deepMerge(src = {}, target = {}) {
  let key;
  for (key in target) {
    src[key] = isObject(src[key]) ? deepMerge(src[key], target[key]) : src[key] = target[key];
  }
  return src;
}
export {
  isString as a,
  isObject as b,
  deepMerge as d,
  goodTimeText as g,
  isFunction as i,
  setObjToUrlParams as s
};
