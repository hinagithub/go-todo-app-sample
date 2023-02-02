
// addボタン
const addBtn = document.getElementById('add-btn')

// closeボタン
const closeBtn = document.getElementById('close-btn')

// sendボタン
const saveBtn = document.getElementById('save-bnt')

// overlay領域
const overlay = document.querySelector('.overlay')

// modal領域
const modal = document.querySelector('.modal')

// モーダルを開く関数
const showModal = () => {
  console.log("クリック");
  overlay.classList.add('on');
  modal.classList.add('on');
}

// モーダルを閉じる関数
const closeModal = () => {
  overlay.classList.remove('on');
  modal.classList.remove('on');
}

// 新規登録関数
const create = ()=>{
  console.log(saveBtn)
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
    console.log("create成功")
    closeModal()
    location.reload()
  })
  .catch((error) => {
    console.log("create失敗")
    closeModal()  
  });
}

// 追加ボタンがクリックされた時
addBtn.addEventListener('click', showModal)

// 閉じるボタンがクリックされた時
closeBtn.addEventListener('click', closeModal)

// モーダル枠外がクリックされた時
modal.addEventListener('click', (event) => {
  if(event.target.closest('.modal-dialog') === null) closeModal()
});

// 保存ボタンがクリックされた時
saveBtn.addEventListener('click', create)

