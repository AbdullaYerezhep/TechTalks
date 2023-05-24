
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
console.log(isAuthenticated);

if (isAuthenticated){
    const myPostsButton = document.querySelector(".headerMyPosts")
    
    myPostsButton.addEventListener("click", () => {
        filterMyPosts()
    })
    
    const filterMyPosts = () => {
        posts.forEach(post => {
            let authorID = post.querySelector(".author-id").getAttribute("author");
            if (authorID !== userID) {
                post.style.display = "none"
            }else{
                post.style.display = "flex"
            }
        })
        
    }
}

// categories filter 

let checkBoxContainers = document.querySelectorAll(".menu-checkbox-container");

checkBoxContainers.forEach(checkBoxContainer => {
  checkBoxContainer.addEventListener("click", () => {
    const checkBoxInput = checkBoxContainer.querySelector(".menuCategoryInput")
    checkBoxContainer.classList.toggle('checked');
    checkBoxInput.checked = !checkBoxInput.checked
    filterByCategory();
  });
});

const filterByCategory = () => {
  let checkedCategories = [];
  let checkedCheckboxes = document.querySelectorAll(".menuCategoryInput:checked");
  
  checkedCheckboxes.forEach(checkbox => {
    let category = checkbox.value;
    checkedCategories.push(category);
  });
  
  
  posts.forEach(post => {
    let postCategories = Array.from(post.querySelectorAll(".post-category"))
    let filteredPosts = postCategories.map(category => category.textContent);
    
    if (checkedCategories.length === 0 || checkedCategories.every(category => filteredPosts.includes(category))) {
      post.style.display = "flex";
    } else {
      post.style.display = "none";
    }
  });
};