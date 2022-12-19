
// function updateState(data) {
//   setState({
//     ...state,
//     Avatar: data.Avatar,
//     UserName: data.UserName,
//     UserDesc: data.UserDesc,
//     Bilibili: data.Bilibili,
//     zhihu: data.zhihu,
//     // Categorys:data.Categorys,
//   })
// };


let url = 'http://139.186.213.52:8082/'
fetch(url, {
  method: 'GET',
}).then(response => response.json()).then(data => {
  if (data === null) {
    window.alert("根据此id查询不到用户")
    console.log("查询失败")
    return
  }
  // updateState(data)
  console.log("调用了fetch")
  function GetInfo() {
    console.log("调用执行了这个函数"+"GetInfo")
  
    return (
      <div class="personal-box">
        <img class="personal-avatar" src={data.Avatar} alt={data.UserName} />
        <p class="personal-name">{data.UserName}</p>
        <p class="personal-desc">{data.UserDesc}</p>
        <span class="gap-dashed"></span>
        <ul class="personal-action">
          <li>
            <i class="iconfont icon-bilibili-line"></i>
            <a href={data.Bilibili} target="_blank"> Bilibili</a>
          </li>

          <li>
            <i class="iconfont icon-zhihu-circle"></i>
            <a href={data.zhihu} target="_blank"> 知乎</a>
            
          </li>

        </ul>
        <span class="gap-solid"></span>
  
  
        {/* <div class="personal-cteg">
          <h2>分类：</h2>
          {data.Categorys.map((item, idx) => (
            <div>
              <i class="iconfont icon-fenlei-copy"></i>
              <a href={"/c/" + item.Cid}>{item.Name}</a>
            </div>
          ))}
        </div> */}
      </div>
    )
  };
  
  ReactDOM.render(<GetInfo />, document.querySelector("#app"));
})
// function GetInfo() {
//   console.log("调用执行了这个函数"+"GetInfo")

//   const [state, setState] = React.useState({
//     Avatar: "/resource/images/1.jpg",
//     UserName: "不会写代码的建筑师",
//     UserDesc: "写代码是医生的邀请",
//     Bilibili: "https://space.bilibili.com/66628849",
//     zhihu: "https://www.zhihu.com/people/xu-heng-wei-63",
//     Categorys: [
//       {
//         "Cid": "1",
//         "Name": "java"
//       },
//       {
//         "Cid": "2",
//         "Name": "golang"
//       },
//     ],
//   });
  
//   return (
//     <div class="personal-box">
//       <img class="personal-avatar" src={state.Avatar} alt={state.UserName} />
//       <h1 class="personal-name">{state.UserName}</h1>
//       <p class="personal-desc">{state.UserDesc}</p>
//       <span class="gap-dashed"></span>
//       <ul class="personal-action">
//         <li>
//           <i class="iconfont icon-bilibili-line"></i>
//           <a href={state.Bilibili} target="_blank"> Bilibili</a>
//         </li>
//         <li>
//           <i class="iconfont icon-zhihu-circle"></i>
//           <a href={state.zhihu} target="_blank"> 知乎</a>
//         </li>
//       </ul>
//       <span class="gap-solid"></span>


//       <div class="personal-cteg">
//         <h2>分类：</h2>
//         {state.Categorys.map((item, idx) => (
//           <div>
//             <i class="iconfont icon-fenlei-copy"></i>
//             <a href={"/c/" + item.Cid}>{item.Name}</a>
//           </div>
//         ))}
//       </div>
//     </div>
//   )
// };

// ReactDOM.render(<GetInfo />, document.querySelector("#app"));
