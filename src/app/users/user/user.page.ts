import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ModalController, NavParams, AlertController } from '@ionic/angular';

import { User } from '../user';
import { URL } from '../../api/init'
import { ObjPage } from '../../api/obj.page';
import { UserService } from '../user.service';

@Component({
  selector: 'app-user',
  templateUrl: './user.page.html',
  styleUrls: ['./user.page.scss'],
})
export class UserPage extends ObjPage<User>{
  constructor(
    private userService: UserService,
    private alertCtrl: AlertController,
    private http: HttpClient,
    protected modalCtrl: ModalController,
    protected navParams: NavParams
  ) {
    super(modalCtrl, navParams)
  }
  get service() { return this.userService }

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
            this.http.patch<User>(`${URL}/api/users/${this.obj.id}/pass?pass=${v['pass']}`, this.obj)
              .subscribe((u: User) => console.log(u))
          }
        }
      ]
    })
    await alert.present();
  }
}
