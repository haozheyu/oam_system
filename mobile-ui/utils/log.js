export function warn(msg) {
  console.log(`%c%s`, 'color: orange', msg)
}

export function info(msg) {
  console.log(`%c%s`, 'color: green', msg)
}

export function error(msg) {
  console.log(`%c%s`, 'color: red', msg)
}

export default {
  warn,
  info,
  error
}
