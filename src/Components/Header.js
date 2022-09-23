import './Page.css';
function header_component() {
    return (
      <div class="header">
          <div class="navleft">
              <img src="./T_Hex.png" class="logo"/>
              <h2>T-Hex</h2>
          </div>
          <ul>
              <li><a href="#">Products</a></li>
              <li><a href="#">Developers</a></li>
              <li><a href="#">Live for Teams</a></li>
              <li><a href="#">Pricing</a></li>
              <li><a href="#">Sign in</a></li>
              <li><input type="button" value="FREE TRAIL" class="btn"/></li>
              <li><i class="fa-solid fa-magnifying-glass"></i></li>
  
          </ul>
      </div>
     
    );
  }
  
  export default header_component;
  