const KEY='birrwise:data:v1'
export const storageService={
  load(){try{return JSON.parse(localStorage.getItem(KEY)||'null')}catch{return null}},
  save(data){localStorage.setItem(KEY,JSON.stringify(data))},
  clear(){localStorage.removeItem(KEY)}
}
