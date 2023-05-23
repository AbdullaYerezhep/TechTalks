
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


// my posts logic

const myPostsButton = document.querySelector(".headerMyPosts")

myPostsButton.addEventListener("click", () => {
    filterMyPosts()
})

const filterMyPosts = () => {

}

// categories filter 

let checkBoxContainers = document.querySelectorAll(".menu-checkbox-container")

checkBoxContainers.forEach(checkBoxContainer => {
    checkBoxContainer.addEventListener("click", ()=>{
        let checkBoxInput = checkBoxContainer.querySelector(".menuCategoryInput")
        checkBoxInput.checked = !checkBoxInput.checked
        checkBoxContainer.classList.toggle('checked');
    })
});

// function checkBox(container) {
//     let checkBoxInput = container.querySelector(".menuCategoryInput")
//     checkBoxInput.checked = !checkBoxInput.checked
//     checkBoxContainer.classList.toggle('checked');
// }

const filterByCategory = () => {
    
}