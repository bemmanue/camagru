$(document).ready(function() {
    let comment_notify = document.getElementById("comment-notify")
    let like_notify = document.getElementById("like-notify")

    comment_notify.addEventListener("change", function () {
        console.log($(comment_notify).val())
    })

    like_notify.addEventListener("change", function () {
        console.log($(like_notify).val())
    })
})