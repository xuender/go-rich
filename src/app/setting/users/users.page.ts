import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ModalController } from '@ionic/angular';
import { User } from '../../api/user';
import { Observable } from 'rxjs';
import { URL } from '../../api/init'
import { UserPage } from './user/user.page';

@Component({
  selector: 'app-users',
  templateUrl: './users.page.html',
  styleUrls: ['./users.page.scss'],
})
export class UsersPage implements OnInit {

  users: Observable<User[]>
  constructor(
    private modalCtrl: ModalController,
    private http: HttpClient
  ) {
    this.load()
  }
  load() {
    this.users = this.http.get<User[]>(`${URL}/api/users`)
  }

  ngOnInit() {
  }

  cancel() {
    this.modalCtrl.dismiss();
  }
  async update(user: User) {
    const modal = await this.modalCtrl.create({
      component: UserPage,
      componentProps: {
        user: user
      }
    });
    modal.onDidDismiss().then(d => {
      if (d.data) {
        this.http.put<User>(`${URL}/api/users/${user.id}`, d.data)
          .subscribe(() => {
            this.load()
          })
      }
    })
    return await modal.present()
  }
  async create() {
    const modal = await this.modalCtrl.create({
      component: UserPage,
      componentProps: {
        user: {
          nick: '新用户',
          extend: {}
        }
      }
    });
    modal.onDidDismiss().then(d => {
      if (d.data) {
        this.http.post<User>(`${URL}/api/users`, d.data)
          .subscribe(() => {
            this.load()
          })
      }
    })
    return await modal.present()
  }
}
