export function setItemToLS(key, value) {
  localStorage.setItem(key, JSON.stringify(value));
}

export function getItemFormLS(key) {
  const value = localStorage.getItem(key);
  if (value) {
    return JSON.parse(value);
  } else {
    return null;
  }
}

export function removeItemInLS(key) {
  localStorage.removeItem(key);
}
