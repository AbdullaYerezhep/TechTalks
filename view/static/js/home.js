
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

if (isAuthenticated){
  const myPostsButton = document.querySelector(".headerMyPosts")
  const headerDetails = document.querySelector(".header-details")
  myPostsButton.addEventListener("click", (e) => {
    headerDetails.open = false;
      e.preventDefault();
      fetch("/", {
        method:"GET",
        headers : {
          "Content-Type": "application/json"
        }
      })
  });
}

// // my posts logic
// console.log(isAuthenticated);


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