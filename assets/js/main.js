
// addボタン
const addBtn = document.getElementById('add-btn')

// closeボタン
const closeBtn = document.getElementById('close-btn')

// sendボタン
const saveBtn = document.getElementById('save-bnt')

// completeボタン
const completeBtn = document.getElementById('complete-btn')

// overlay領域
const overlay = document.querySelector('.overlay')

// modal領域
const modal = document.querySelector('.modal')

// modalタイトル
const modalTitle = document.getElementById("modal-title")


// 新規モーダルを開く関数
const showCreateModal = () => {
  overlay.classList.add('on');
  modal.classList.add('on');
  modalTitle.innerHTML='新規作成'
  const todoId = document.getElementById('todo-id')
  todoId.value=null

}

// 編集モーダルを開く関数
const showEditModal = (id) => {
  overlay.classList.add('on');
  modal.classList.add('on');
  modalTitle.innerHTML='編集 ID:' + id
  const todoId = document.getElementById('todo-id')
  todoId.value=id
}


// モーダルを閉じる関数
const closeModal = () => {
  overlay.classList.remove('on');
  modal.classList.remove('on');
  modalTitle.innerHTML=''
}

// 新規登録関数
const create = ()=>{
  console.log("保存ボタンクリック")
  const text = document.getElementById("todo-text").value
  console.log('text: ', text)
  fetch("http://localhost:3000/todo",{
    method: "POST",
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({body:text})
  })
  .then((response) => {
    closeModal()
    location.reload()
  })
  .catch((error) => {
    console.log("create失敗")
    closeModal()  
  });
}

// 編集関数
const update = (id)=>{  
  const text = document.getElementById("todo-text").value
  fetch("http://localhost:3000/todo/" + id,{
    method: "POST",
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      completed: false,
      body:text,
    })
  })
  .then((response) => {
    closeModal()
    location.reload()
  })
  .catch((error) => {
    console.log("update失敗")
    closeModal()  
  });
}

// 完了済みにする
const complete = (id, text)=>{
  console.log(id, TextDecoderStream)
  fetch("http://localhost:3000/todo/" + id,{
    method: "POST",
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      completed: true,
      body:text,
    })
  })
  .then((response) => {
    closeModal()
    location.reload()
  })
  .catch((error) => {
    console.log("update失敗")
    closeModal()  
  });
}

const save =()=>{
  const todoId = document.getElementById('todo-id').value
  if(todoId) update(todoId)
  else create()
}

// モーダル枠外がクリックされた時
modal.addEventListener('click', (event) => {
  if(event.target.closest('.modal-dialog') === null) closeModal()
})
