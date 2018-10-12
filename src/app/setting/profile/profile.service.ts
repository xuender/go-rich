import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { AlertController } from '@ionic/angular';
import { Router } from '@angular/router';

import { Config } from '../../api/config'
import { User } from '../../users/user';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  constructor(
    private http: HttpClient,
    private alertCtrl: AlertController,
    private router: Router,
  ) { }

  get profile$() {
    return this.http.get<any>(`${Config.URL}/api/profile`)
  }

  save(user: User) {
    return this.http.patch<User>(`${Config.URL}/api/profile`, user)
  }

  logout() {
    localStorage.removeItem('token')
  }

  login() {
    this.router.navigateByUrl('/start')
  }

  error(msg: string) {
    this.alertCtrl.create({
      header: '错误',
      message: msg,
      buttons: ['确定']
    }).then(a => a.present())
  }

  async pass() {
    const alert = await this.alertCtrl.create({
      header: '修改密码',
      subHeader: '修改当前用户登录密码',
      inputs: [
        {
          name: 'old',
          label: '旧密码',
          type: 'password',
          placeholder: '用户旧密码'
        },
        {
          name: 'pass',
          label: '新密码',
          type: 'password',
          placeholder: '用户新密码'
        },
        {
          name: 'two',
          label: '重复新密码',
          type: 'password',
          placeholder: '重复新密码'
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
            if (v['pass'] != v['two']) { return false }
            this.http.patch<string>(`${Config.URL}/api/profile/pass?old=${v['old']}&pass=${v['pass']}`, {})
              .subscribe((t: string) => localStorage.setItem('token', t))
            return true
          }
        }
      ]
    })
    await alert.present();
  }
}
