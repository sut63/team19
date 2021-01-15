import React, { FC,useState } from 'react';
import * as plugins from './plugins';
import { createApp, SidebarPage } from '@backstage/core';
import { AppSidebar } from './sidebar';
import Login from './components/Login/Login';
import Sign from './components/Sign-in/Sign-in'

const App: FC<{}> = () => {

  const app = createApp({
    plugins: Object.values(plugins),
  });
  
  const AppProvider = app.getProvider();
  const AppRouter = app.getRouter();
  const AppRoutes = app.getRoutes();
  
  const [Session,setSession]  = useState({
    isLoggedIn: false,
    isSignIn: false,
    currentUser: null,
    errorMessage: null
  
  });

  if(Session.isLoggedIn){

  }

  if(Session.isSignIn){

  }

return(
  <AppProvider>
      <AppRouter>
        <SidebarPage>
            {Session.isSignIn ? (
             <header>
              {Session.isLoggedIn ? (
                <header>
                  <AppRoutes/>
                  <AppSidebar/>
                </header>
              ) : (
              <Sign setSession = {setSession}/>
              )}
             </header>
            ) : (
              <Login setSession = {setSession}/>
            )}
        </SidebarPage>     
    </AppRouter>
  </AppProvider>
  );
};

export default App;