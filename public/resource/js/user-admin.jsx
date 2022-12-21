function UserComName(props){
    return(
        <div class="row mb-3">
            <label for="inputEmail3" class="col-sm-2 col-form-label">{props.comName}：</label>
            <div class="col-sm-10">
            <input type="text" class="form-control" id="inputText"></input>
            </div>
        </div>
    );
}
function UserComMail(){
    return(
        <div class="row mb-3">
            <label for="inputEmail3" class="col-sm-2 col-form-label">Email：</label>
            <div class="col-sm-10">
            <input type="email" class="form-control" id="inputEmail"></input>
            </div>
        </div>
    );
}
function UserComPassword(){
    return(
        <div class="row mb-3">
            <label for="inputPassword3" class="col-sm-2 col-form-label">Password</label>
            <div class="col-sm-10">
            <input type="password" class="form-control" id="inputPassword"></input>
            </div>
        </div>
    );
}
function UserComSpan(props){
    return(
        <div class="row mb-3">
        <label for="inputEmail3" class="col-sm-2 col-form-label">{props.name}：</label>
        <div class="col-sm-10" style={{"align-items":"center","display":"flex"}}>
        <span>{props.value}</span>
        </div>
    </div>
    )

}

function UserInfo() {
    return (
        <div>
        <main class="main" id="main">  

        <div class="pagetitle">
          <h1>个人信息</h1>
        </div>
        <section class="section">
          <div class="row">
              <div class="card">
                <div class="card-body" style={{"padding-top":"50px"}}>
                  <form >
                  <UserComName  comName="用户昵称"/>
                  <UserComSpan name="用户id" value="1258413"/>
                  <UserComMail  />
                  <UserComPassword />
                  <UserComName  comName="知乎"/>
                  <UserComName  comName="B站"/>
                  <UserComName  comName="Github"/>

                    <fieldset class="row mb-3">
                      <legend class="col-form-label col-sm-2 pt-0">性别</legend>
                      <div class="col-sm-10">
                        <div class="form-check">
                          <input class="form-check-input" type="radio" name="gridRadios" id="gridRadios1" value="option1" checked></input>
                          <label class="form-check-label" for="gridRadios1">
                            男
                          </label>
                        </div>
                        <div class="form-check">
                          <input class="form-check-input" type="radio" name="gridRadios" id="gridRadios2" value="option2"></input>
                          <label class="form-check-label" for="gridRadios2">
                            女
                          </label>
                        </div>
                        <div class="form-check">
                          <input class="form-check-input" type="radio" name="gridRadios" id="gridRadios3" value="option3"></input>
                          <label class="form-check-label" for="gridRadios3">
                            保密
                          </label>
                        </div>
                      </div>
                    </fieldset>
                    <div class="text-center">
                      <button type="submit" class="btn btn-primary">Submit</button>
                    </div>
                  </form>
    
                </div>
              </div>
          </div>
        </section>
        </main>
        </div>
    );
}
function SecondApp() {
    return (
        <div class="main" id="main"><h1>第er类型</h1></div>
    );
}
function ThridApp() {
    return (
        <div class="main" id="main"><h1>第san类型</h1></div>
    );
}


function Click(props,event) {
    if (props === "个人信息") {
        ReactDOM.render(<UserInfo />, document.querySelector("#app"));
    } else if (props === "安全设置") {
        ReactDOM.render(<SecondApp />, document.querySelector("#app"));
    } else {
        return <UserInfo />;
    }   
}
function NavSelect() {
    return (
        <div>
        <li class="nav-item" >
                <a class="nav-link collapsed" onClick={Click.bind(this, "个人信息")} >
                <i class="bi bi-grid"></i>
                <span>个人信息</span>
                </a>
            </li>
            <li class="nav-item">
                <a class="nav-link collapsed" onClick={Click.bind(this, "安全设置")}>
                <i class="bi bi-grid"></i>
                <span>安全设置</span>
                </a>
            </li>
            <li class="nav-item">
              <a className={ "nav-link collapsed"}  onClick={Click.bind(this, "Dashboard")}>
              <i class="bi bi-grid"></i>
              <span>Dashboard</span>
              </a>
            </li>
            </div>
    );
}


var TOKEN_KEY = "AUTH_TOKEN";
var USER_KEY = "USER_INFO";
function test(){
  if (localStorage.getItem(TOKEN_KEY)) {
    console.log(localStorage.getItem(USER_KEY))
  }else{
    console.log("没有找到token")
    window.location.replace("/login");
    console.log("退出登录");
  }
}
test()



ReactDOM.render(<UserInfo />, document.querySelector("#app"));
ReactDOM.render(<NavSelect />, document.querySelector("#navSelect"));