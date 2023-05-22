
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

const myPostsButton = document.getElementsByClassName("headerMyPosts")

myPostsButton.addEventListener("click", () => {
    
})