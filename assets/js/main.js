console.log('main.js called')

// addボタン
const addBtn = document.getElementById('add-btn')

// モーダル
const modal = document.getElementById('modal')

// モーダルを開く関数
const openModal = () => {
  modal.className = 'modal-open'
}

// Addボタンがクリックされた時
addBtn.addEventListener('click', openModal)
