import { NgModule  } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouteReuseStrategy } from '@angular/router';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';


import { IonicModule, IonicRouteStrategy } from '@ionic/angular';
import { SplashScreen } from '@ionic-native/splash-screen/ngx';
import { StatusBar } from '@ionic-native/status-bar/ngx';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ExtPageModule } from './ext/ext.module';
import { XlsxesPageModule } from './setting/xlsxes/xlsxes.module';
import { FormsModule } from '@angular/forms';
import { ErrorInterceptor } from './api/error.interceptor';
import { JWTInterceptor } from './api/jwt.interceptor';
import { ProfilePageModule } from './setting/profile/profile.module';
import { TagsPageModule } from './tags/tags.module';
import { SelectPageModule } from './tags/select/select.module';
import { UsersPageModule } from './users/users.module';

@NgModule({
  declarations: [
    AppComponent,
  ],
  entryComponents: [
  ],
  imports: [
    HttpClientModule,
    BrowserModule,
    IonicModule.forRoot(),
    AppRoutingModule,
    // service 扩展
    ExtPageModule,
    XlsxesPageModule,
    UsersPageModule,
    ProfilePageModule,
    FormsModule,
    TagsPageModule,
    SelectPageModule,
  ],
  providers: [
    StatusBar,
    SplashScreen,
    { provide: RouteReuseStrategy, useClass: IonicRouteStrategy },
    { provide: HTTP_INTERCEPTORS, useClass: ErrorInterceptor, multi: true },
    { provide: HTTP_INTERCEPTORS, useClass: JWTInterceptor, multi: true },
  ],
  bootstrap: [
    AppComponent,
  ],
})
export class AppModule { }
