import { useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { ArrowLeft } from 'lucide-react'
import { useApp } from '../context/AppContext'
import { paymentMethods } from '../data/defaults'
import { validateTransaction } from '../utils/finance'

const blank={title:'',amount:'',type:'expense',categoryId:'food',date:new Date().toISOString().slice(0,10),paymentMethod:'Cash',description:''}
export default function TransactionForm(){const {id}=useParams();const nav=useNavigate();const {transactions,categories,addTransaction,updateTransaction,t}=useApp();const existing=id?transactions.find(x=>x.id===id):null;const [form,setForm]=useState(existing||blank);const [error,setError]=useState('')
 useEffect(()=>{if(existing)setForm(existing)},[existing])
 const set=(key,value)=>setForm(v=>({...v,[key]:value}))
 const submit=e=>{e.preventDefault();const err=validateTransaction(form);if(err){setError(err);return}const item={...form,amount:Number(form.amount)};if(id)updateTransaction(id,item);else addTransaction(item);nav('/transactions')}
 const relevant=categories.filter(c=>c.type===form.type)
 return <div><button className="back-btn" onClick={()=>nav(-1)}><ArrowLeft size={18}/> {t('back')}</button><div className="page-heading"><div><p>{t('transactionDetails')}</p><h1>{id?t('edit'):t('addTransaction')}</h1></div></div>
  <form className="form-card" onSubmit={submit}><div className="type-toggle"><button type="button" className={form.type==='expense'?'active expense':''} onClick={()=>{set('type','expense');set('categoryId','food')}}>{t('expense')}</button><button type="button" className={form.type==='income'?'active income':''} onClick={()=>{set('type','income');set('categoryId','salary')}}>{t('income')}</button></div>{error&&<div className="form-error">{error}</div>}
  <div className="form-grid"><label>{t('title')}<input required value={form.title} onChange={e=>set('title',e.target.value)} placeholder={t('titlePlaceholder')}/></label><label>{t('amountEtb')}<input required min="0.01" step="0.01" type="number" value={form.amount} onChange={e=>set('amount',e.target.value)} placeholder="0.00"/></label><label>{t('category')}<select value={form.categoryId} onChange={e=>set('categoryId',e.target.value)}>{relevant.map(c=><option value={c.id} key={c.id}>{c.name}</option>)}</select></label><label>{t('date')}<input required type="date" max={new Date().toISOString().slice(0,10)} value={form.date} onChange={e=>set('date',e.target.value)}/></label><label>{t('paymentMethod')}<select value={form.paymentMethod} onChange={e=>set('paymentMethod',e.target.value)}>{paymentMethods.map(x=><option key={x}>{x}</option>)}</select></label><label className="full">{t('description')}<textarea rows="4" value={form.description||''} onChange={e=>set('description',e.target.value)} placeholder={t('optionalNote')}/></label></div><div className="form-actions"><button type="button" className="secondary-btn" onClick={()=>nav(-1)}>{t('cancel')}</button><button className="primary-btn" type="submit">{t('save')}</button></div></form></div>
}
