import { Injectable } from '@angular/core';
import { AlertController, NavController } from '@ionic/angular';
import { HttpClient } from '@angular/common/http';

import { URL } from './init'

@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  constructor(
    private http: HttpClient,
    private alertCtrl: AlertController,
    private navCtrl: NavController,
  ) { }
  logout() {
    localStorage.setItem('token', '')
  }
  login() {
    this.alertCtrl.create({
      header: '用户登录',
      subHeader: '请输入登录名和密码',
      inputs: [
        {
          name: 'nick',
          label: '登录名',
          type: 'text',
          placeholder: '姓名或手机号'
        },
        {
          name: 'pass',
          label: '密码',
          type: 'password',
          placeholder: '登录密码'
        },
      ],
      buttons: [
        {
          text: '取消',
          role: 'cancel',
          cssClass: 'secondary',
        }, {
          text: '确定',
          handler: (v) => {
            this.http.get<any>(`${URL}/login?nick=${v['nick']}&pass=${v['pass']}`)
              .subscribe(a => {
                localStorage.setItem('token', a.token)
                this.navCtrl.goBack()
              })
          }
        }
      ]
    }).then(a => a.present());
  }
  error(msg: string) {
    this.alertCtrl.create({
      header: '错误',
      message: msg,
      buttons: ['确定']
    }).then(a => a.present())
  }
}
