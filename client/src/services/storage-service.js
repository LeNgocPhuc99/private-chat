export function setItemToSS(key, value) {
  sessionStorage.setItem(key, JSON.stringify(value));
}

export function getItemFormSS(key) {
  const value = sessionStorage.getItem(key);
  if (value) {
    return JSON.parse(value);
  } else {
    return null;
  }
}

export function removeItemInSS(key) {
  sessionStorage.removeItem(key);
}
