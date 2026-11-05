import { createContext, useContext, useEffect, useMemo, useState } from 'react'
import { defaultCategories } from '../data/defaults'
import { storageService } from '../services/storage'
import { apiService } from '../services/api'
import dictionaries from '../locales'

const AppContext=createContext(null)
const today=new Date()
const iso=(days)=>{const d=new Date(today);d.setDate(d.getDate()+days);return d.toISOString().slice(0,10)}
const seed={
  transactions:[
    {id:'seed-1',title:'Monthly salary',amount:28000,type:'income',categoryId:'salary',date:iso(-2),paymentMethod:'Bank transfer',description:'Monthly employment income'},
    {id:'seed-2',title:'Mercato groceries',amount:1850,type:'expense',categoryId:'food',date:iso(-1),paymentMethod:'Telebirr',description:''},
    {id:'seed-3',title:'Apartment rent',amount:8500,type:'expense',categoryId:'housing',date:iso(-12),paymentMethod:'Bank transfer',description:''},
    {id:'seed-4',title:'Taxi and fuel',amount:720,type:'expense',categoryId:'transport',date:iso(-6),paymentMethod:'Cash',description:''},
    {id:'seed-5',title:'Freelance project',amount:4200,type:'income',categoryId:'freelance',date:iso(-18),paymentMethod:'CBE Birr',description:''}
  ],categories:defaultCategories,budgets:[{id:'budget-1',categoryId:'food',month:iso(0).slice(0,7),limitAmount:3500},{id:'budget-2',categoryId:'transport',month:iso(0).slice(0,7),limitAmount:2200}],preferences:{currency:'ETB',language:'en',theme:'light'}
}

export function AppProvider({children}){
  const saved=storageService.load()
  const initial=saved&&typeof saved==='object'?saved:seed
  const [transactions,setTransactions]=useState(initial.transactions||seed.transactions)
  const [categories,setCategories]=useState(initial.categories||seed.categories)
  const [budgets,setBudgets]=useState(initial.budgets||seed.budgets)
  const [preferences,setPreferences]=useState(initial.preferences||seed.preferences)
  useEffect(()=>{storageService.save({transactions,categories,budgets,preferences})},[transactions,categories,budgets,preferences])
  useEffect(()=>{document.documentElement.classList.toggle('dark',preferences.theme==='dark')},[preferences.theme])
  useEffect(()=>{
    Promise.all([apiService.listTransactions(),apiService.listCategories(),apiService.listBudgets(),apiService.getPreferences()])
      .then(([apiTransactions,apiCategories,apiBudgets,apiPreferences])=>{
        setTransactions(apiTransactions)
        setCategories(apiCategories)
        setBudgets(apiBudgets)
        setPreferences(apiPreferences)
      })
      .catch(()=>{})
  },[])
  useEffect(()=>{apiService.updatePreferences(preferences).catch(()=>{})},[preferences])
  const t=(key)=>dictionaries[preferences.language]?.[key]||dictionaries.en[key]||key
  const categoryById=useMemo(()=>Object.fromEntries(categories.map(c=>[c.id,c])),[categories])
  const addTransaction=async(item)=>{try{const created=await apiService.createTransaction(item);setTransactions(v=>[created,...v])}catch{setTransactions(v=>[{...item,id:crypto.randomUUID(),createdAt:new Date().toISOString()},...v])}}
  const updateTransaction=async(id,item)=>{try{const updated=await apiService.updateTransaction(id,item);setTransactions(v=>v.map(t=>t.id===id?updated:t))}catch{setTransactions(v=>v.map(t=>t.id===id?{...t,...item,updatedAt:new Date().toISOString()}:t))}}
  const deleteTransaction=async(id)=>{try{await apiService.deleteTransaction(id)}catch{};setTransactions(v=>v.filter(t=>t.id!==id))}
  const updateCategories=(updater)=>setCategories(old=>{const next=typeof updater==='function'?updater(old):updater;const oldById=new Map(old.map(x=>[x.id,x]));next.forEach(item=>{const existing=oldById.get(item.id);if(!existing){apiService.createCategory(item).then(created=>setCategories(v=>v.map(x=>x.id===item.id?created:x))).catch(()=>{})}else if(JSON.stringify(existing)!==JSON.stringify(item)){apiService.updateCategory(item.id,item).catch(()=>{})}});old.forEach(item=>{if(!next.some(x=>x.id===item.id)){apiService.deleteCategory(item.id).catch(()=>{})}});return next})
  const updateBudgets=(updater)=>setBudgets(old=>{const next=typeof updater==='function'?updater(old):updater;const oldById=new Map(old.map(x=>[x.id,x]));next.forEach(item=>{const existing=oldById.get(item.id);if(!existing){apiService.createBudget(item).then(created=>setBudgets(v=>v.map(x=>x.id===item.id?created:x))).catch(()=>{})}else if(JSON.stringify(existing)!==JSON.stringify(item)){apiService.updateBudget(item.id,item).catch(()=>{})}});old.forEach(item=>{if(!next.some(x=>x.id===item.id)){apiService.deleteBudget(item.id).catch(()=>{})}});return next})
  const reset=()=>{storageService.clear();setTransactions(seed.transactions);setCategories(seed.categories);setBudgets(seed.budgets);setPreferences(seed.preferences)}
  return <AppContext.Provider value={{transactions,categories,budgets,preferences,setPreferences,t,categoryById,addTransaction,updateTransaction,deleteTransaction,setCategories:updateCategories,setBudgets:updateBudgets,reset}}>{children}</AppContext.Provider>
}
export const useApp=()=>useContext(AppContext)
