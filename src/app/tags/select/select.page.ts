import { Component } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';
import { Observable } from 'rxjs';
import { pull, includes } from 'lodash'

import { Tag } from '../tag';

@Component({
  selector: 'app-select',
  templateUrl: './select.page.html',
  styleUrls: ['./select.page.scss'],
})
export class SelectPage {
  tags$: Observable<Tag[]>
  obj: any = {}
  constructor(
    private modalCtrl: ModalController,
    private navParams: NavParams,
  ) {
    this.tags$ = this.navParams.get('tags$')
    const tags = this.navParams.get('tags')
    for (const t of tags) {
      this.obj[t] = true
    }
  }

  cancel() {
    this.modalCtrl.dismiss();
  }

  select(tag: string = '') {
    const tags: string[] = []
    for (const k in this.obj) {
      if (this.obj[k]) {
        tags.push(k)
      }
    }
    if (tag && !includes(tags, tag)) {
      tags.push(tag)
    }
    this.modalCtrl.dismiss(tags);
  }

  change(tag: string) {
    if (tag in this.obj) {
      this.obj[tag] = !this.obj[tag]
    } else {
      this.obj[tag] = true
    }
  }
}
