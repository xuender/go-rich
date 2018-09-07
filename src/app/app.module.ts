import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule, RouteReuseStrategy } from '@angular/router';
import { HttpClientModule } from '@angular/common/http';


import { IonicModule, IonicRouteStrategy } from '@ionic/angular';
import { SplashScreen } from '@ionic-native/splash-screen/ngx';
import { StatusBar } from '@ionic-native/status-bar/ngx';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { SettingPageModule } from './setting/setting.module';
import { ExtPageModule } from './setting/ext/ext.module';
import { CustomerPageModule } from './contact/customer/customer.module';
import { XlsxesPageModule } from './setting/xlsxes/xlsxes.module';
import { UsersPageModule } from './setting/users/users.module';
import { HomePageModule } from './home/home.module';
import { ContactPageModule } from './contact/contact.module';

@NgModule({
  declarations: [
    AppComponent,
  ],
  entryComponents: [],
  imports: [
    HttpClientModule,
    BrowserModule,
    IonicModule.forRoot(),
    AppRoutingModule,
    // service 扩展
    ExtPageModule,
    XlsxesPageModule,
    UsersPageModule,
  ],
  providers: [
    StatusBar,
    SplashScreen,
    { provide: RouteReuseStrategy, useClass: IonicRouteStrategy }
  ],
  bootstrap: [
    AppComponent,
  ]
})
export class AppModule { }
