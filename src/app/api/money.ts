export function ToMoney(obj: any, end = 'Money') {
  return new Proxy(obj, {
    get: (target: any, p: string) => {
      if (p.endsWith(end)) {
        return target[p.substr(0, p.indexOf(end))] / 100
      } else {
        return target[p]
      }
    },
    set: (target: any, p: string, value: any) => {
      if (p.endsWith('Money')) {
        target[p.substr(0, p.indexOf('Money'))] = value * 100
      } else {
        return target[p] = value
      }
    }
  })
}
