import { Component, OnInit } from '@angular/core';
import { ModalController, NavParams, AlertController } from '@ionic/angular';
import { HttpClient } from '@angular/common/http';
import { URL } from '../../../api/init'
import { User } from '../../../api/user';

@Component({
  selector: 'app-user',
  templateUrl: './user.page.html',
  styleUrls: ['./user.page.scss'],
})
export class UserPage implements OnInit {
  user: User
  constructor(
    private alertCtrl: AlertController,
    private http: HttpClient,
    private modalCtrl: ModalController,
    private navParams: NavParams
  ) {
    this.user = this.navParams.get('user')
  }

  ngOnInit() {
  }

  cancel() {
    this.modalCtrl.dismiss();
  }

  ok() {
    this.modalCtrl.dismiss(this.user);
  }
  async pass() {

    const alert = await this.alertCtrl.create({
      header: '密码重置',
      subHeader: '重新设置用户登录密码',
      inputs: [
        {
          name: 'pass',
          label: '新密码',
          type: 'password',
          placeholder: '用户新密码'
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
            // console.log(v)
            this.http.patch<User>(`${URL}/api/users/${this.user.id}/pass?pass=${v['pass']}`, this.user)
              .subscribe((u: User) => console.log(u))
          }
        }
      ]
    });

    await alert.present();
  }
}
