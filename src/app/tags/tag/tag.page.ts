import { Component } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';

import { Tag } from '../tag';
import { TagService } from '../tag.service';
import { ObjPage } from '../../api/obj.page';

@Component({
  selector: 'app-tag',
  templateUrl: './tag.page.html',
  styleUrls: ['./tag.page.scss'],
})
export class TagPage extends ObjPage<Tag>{
  constructor(
    public tagService: TagService,
    protected modalCtrl: ModalController,
    protected navParams: NavParams
  ) {
    super(modalCtrl, navParams)
    if (!this.obj.use) {
      this.obj.use = {}
    }
  }
  get service() { return this.tagService }
}
