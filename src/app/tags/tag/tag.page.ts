import { Component } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';

import { Tag } from '../tag';
import { TagService } from '../tag.service';

@Component({
  selector: 'app-tag',
  templateUrl: './tag.page.html',
  styleUrls: ['./tag.page.scss'],
})
export class TagPage {
  tag: Tag
  constructor(
    public tagService: TagService,
    private modalCtrl: ModalController,
    private navParams: NavParams
  ) {
    this.tag = this.navParams.get('tag')
    if (!this.tag.use) {
      this.tag.use = {}
    }
  }

  cancel() {
    this.modalCtrl.dismiss();
  }

  async save() {
    this.modalCtrl.dismiss(await this.tagService.save(this.tag))
  }
}
