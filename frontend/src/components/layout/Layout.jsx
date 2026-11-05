import { NavLink, Outlet, useNavigate } from 'react-router-dom'
import { LayoutDashboard, ArrowLeftRight, Plus, Tags, WalletCards, ChartNoAxesCombined, Settings, Info, Moon, Sun, Menu, X } from 'lucide-react'
import { useState } from 'react'
import { useApp } from '../../context/AppContext'

export default function Layout(){
 const {t,preferences,setPreferences}=useApp();const [open,setOpen]=useState(false);const nav=useNavigate()
 const links=[['/dashboard','dashboard',LayoutDashboard],['/transactions','transactions',ArrowLeftRight],['/categories','categories',Tags],['/budgets','budgets',WalletCards],['/reports','reports',ChartNoAxesCombined],['/settings','settings',Settings],['/about','about',Info]]
 const toggleTheme=()=>setPreferences(p=>({...p,theme:p.theme==='light'?'dark':'light'}))
 return <div className="app-shell">
  <aside className={`sidebar ${open?'open':''}`}><div className="brand"><div className="brand-mark">B</div><div><strong>BirrWise</strong><span>Finance tracker</span></div><button className="icon-btn sidebar-close" onClick={()=>setOpen(false)}><X size={20}/></button></div>
   <nav>{links.map(([path,key,Icon])=>{const active=window.location.pathname===path;return <NavLink key={path} to={path} className={({isActive})=>`nav-link ${isActive?'active':''}`} onClick={()=>setOpen(false)}><Icon size={19}/><span>{t(key)}</span></NavLink>})}</nav>
   <button className="theme-row" onClick={toggleTheme}>{preferences.theme==='light'?<Moon size={18}/>:<Sun size={18}/>}<span>{preferences.theme==='light'?'Dark mode':'Light mode'}</span></button>
  </aside>{open&&<div className="overlay" onClick={()=>setOpen(false)}/>}<main className="main"><header className="topbar"><button className="icon-btn menu-btn" onClick={()=>setOpen(true)}><Menu size={22}/></button><div><p>Personal finance</p><strong>Good {new Date().getHours()<12?'morning':new Date().getHours()<18?'afternoon':'evening'}</strong></div><button className="avatar">OG</button></header><div className="content"><Outlet/></div></main>
  <button className="fab" title={t('addTransaction')} onClick={()=>nav('/transactions/new')}><Plus size={24}/></button>
  <nav className="mobile-nav">{links.slice(0,5).map(([path,key,Icon])=><NavLink key={path} to={path} className={({isActive})=>isActive?'active':''}><Icon size={20}/><span>{t(key).split(' ')[0]}</span></NavLink>)}</nav>
 </div>
}
