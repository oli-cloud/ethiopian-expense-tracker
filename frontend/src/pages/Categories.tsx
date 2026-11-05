import { useState } from 'react'
import { Plus, Pencil, Trash2, X } from 'lucide-react'
import { useApp } from '../context/AppContext'

export default function Categories() {
  const { categories, setCategories, categoryById, t } = useApp()
  const [name, setName] = useState('')
  const [type, setType] = useState('expense')
  const [editing, setEditing] = useState(null)
  const [editName, setEditName] = useState('')
  const [editType, setEditType] = useState('expense')

  const createCategory = e => {
    e.preventDefault()
    if (!name.trim()) return
    setCategories(v => [...v, {
      id: crypto.randomUUID(),
      name: name.trim(),
      type,
      color: type === 'income' ? '#10b981' : '#64748b',
      icon: 'Ellipsis'
    }])
    setName('')
  }

  const openEdit = category => {
    setEditing(category)
    setEditName(category.name)
    setEditType(category.type)
  }

  const closeEdit = () => setEditing(null)

  const saveEdit = e => {
    e.preventDefault()
    if (!editing || !editName.trim()) return
    setCategories(v => v.map(category => category.id === editing.id
      ? { ...category, name: editName.trim(), type: editType }
      : category))
    closeEdit()
  }

  return <div>
    <div className="page-heading">
      <div><p>{t('organizeSpending')}</p><h1>{t('categories')}</h1></div>
    </div>
    <div className="two-column">
      <section className="panel">
        <div className="panel-head"><h3>{t('allCategoriesTitle')}</h3></div>
        <div className="category-list">
          {categories.map(category => {
            const localized = categoryById[category.id]?.name || category.name
            return <div className="category-row" key={category.id}>
              <span className="category-dot" style={{ background: category.color }}>{localized.charAt(0)}</span>
              <div>
                <strong>{localized}</strong>
                <small>{category.type === 'income' ? t('income') : t('expense')}</small>
              </div>
              <div className="row-actions">
                <button className="icon-btn" onClick={() => openEdit(category)} title={t('editCategory')}><Pencil size={16} /></button>
                {!category.id.includes('-') && <button className="icon-btn danger" onClick={() => confirm(t('deleteCategory')) && setCategories(v => v.filter(x => x.id !== category.id))} title={t('delete')}><Trash2 size={16} /></button>}
              </div>
            </div>
          })}
        </div>
      </section>

      <form className="panel form-card compact" onSubmit={createCategory}>
        <h3>{t('createCategory')}</h3>
        <label>{t('categoryName')}<input required value={name} onChange={e => setName(e.target.value)} placeholder={t('categoryName')} /></label>
        <label>{t('category')}<select value={type} onChange={e => setType(e.target.value)}><option value="expense">{t('expense')}</option><option value="income">{t('income')}</option></select></label>
        <button className="primary-btn" type="submit"><Plus size={18} />{t('addCategory')}</button>
      </form>
    </div>

    {editing && <div className="modal-backdrop" role="presentation" onMouseDown={e => { if (e.target === e.currentTarget) closeEdit() }}>
      <div className="category-modal" role="dialog" aria-modal="true" aria-labelledby="edit-category-title">
        <div className="modal-head">
          <h3 id="edit-category-title">{t('editCategory')}</h3>
          <button className="icon-btn" type="button" onClick={closeEdit} aria-label={t('cancel')}><X size={18} /></button>
        </div>
        <form onSubmit={saveEdit} className="modal-form">
          <label>{t('categoryName')}<input autoFocus required value={editName} onChange={e => setEditName(e.target.value)} /></label>
          <label>{t('category')}<select value={editType} onChange={e => setEditType(e.target.value)}><option value="expense">{t('expense')}</option><option value="income">{t('income')}</option></select></label>
          <div className="modal-actions">
            <button className="secondary-btn" type="button" onClick={closeEdit}>{t('cancel')}</button>
            <button className="primary-btn" type="submit">{t('updateCategory')}</button>
          </div>
        </form>
      </div>
    </div>}
  </div>
}
