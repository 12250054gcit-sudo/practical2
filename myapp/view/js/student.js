function addStudent(){
    var data = {
        //retrieve from data
        stdid : parseInt(document.getElementById("sid").value),//while storing data in db it should be in int format
        fname : document.getElementById("fname").value,
        lname : document.getElementById("lname").value,
        email : document.getElementById("email").value
    }
    //call the add student api to store from data
    fetch('/student/add',{
        method: "POST",
        body: JSON.stringify(data),//converting js object to json string type
        headers: {"Content-type": "application/json; charset-UTF-8"}
    });
}