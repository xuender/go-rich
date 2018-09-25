import { Component } from '@angular/core';
import { ModalController, ActionSheetController } from '@ionic/angular';
import { Observable } from 'rxjs';

import { User } from './user';
import { ObjsPage } from '../api/objs.page';
import { UserPage } from './user/user.page';
import { UserService } from './user.service';

@Component({
  selector: 'app-users',
  templateUrl: './users.page.html',
  styleUrls: ['./users.page.scss'],
})
export class UsersPage extends ObjsPage<User> {
  users$: Observable<User[]>
  constructor(
    public userService: UserService,
    protected modalCtrl: ModalController,
    protected actionSheetCtrl: ActionSheetController,
  ) {
    super(modalCtrl, actionSheetCtrl)
    this.users$ = this.userService.all$
  }
  get service() { return this.userService }
  get page() { return UserPage }
  get newObj() { return {} }
  get title() { return '用户' }
}
