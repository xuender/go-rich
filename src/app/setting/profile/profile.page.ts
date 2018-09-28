import { ModalController, AlertController, ToastController } from '@ionic/angular';
import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { URL } from '../../api/init'
import { ProfileService } from '../../api/profile.service';
@Component({
  selector: 'app-profile',
  templateUrl: './profile.page.html',
  styleUrls: ['./profile.page.scss'],
})
export class ProfilePage implements OnInit {
  user: any = { id: '', nick: '' }
  constructor(
    private modalCtrl: ModalController,
    private profile: ProfileService,
    private http: HttpClient,
    private alertCtrl: AlertController,
    private toastCtrl: ToastController,
  ) {
    this.http.get<any>(`${URL}/api/profile`)
      .subscribe(p => this.user = p)
  }

  ngOnInit() {
  }
  logout() {
    this.profile.logout()
    this.modalCtrl.dismiss();
    this.profile.login()
  }
  cancel() {
    this.modalCtrl.dismiss();
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
            this.http.patch<string>(`${URL}/api/profile?old=${v['old']}&pass=${v['pass']}`, {})
              .subscribe((t: string) => localStorage.setItem('token', t))
            return true
          }
        }
      ]
    })
    await alert.present();
  }
}
