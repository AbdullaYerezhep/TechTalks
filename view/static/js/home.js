
const posts = document.querySelectorAll(".post")

posts.forEach(post => {
    post.addEventListener("click", () =>{
        openPost(post)
    })
});

function openPost(post) {
    console.log(post);
    let id = post.querySelector(".id").textContent
    window.location.href = "/post/?id="+id
}


// categories filter 

let checkBoxContainers = document.querySelectorAll(".menu-checkbox-container")

checkBoxContainers.forEach(checkBoxContainer => {
    checkBoxContainer.addEventListener("click", ()=>{
        checkBox(checkBoxContainer)
    })
});

function checkBox(container) {
    let checkBoxInput = container.querySelector(".menuCategoryInput")
    checkBoxInput.checked = !checkBoxInput.checked
}