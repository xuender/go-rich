import { Component, OnInit } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';
import { Observable } from 'rxjs';
import { pull, includes } from 'lodash'

import { Tag } from '../tag';

@Component({
  selector: 'app-select',
  templateUrl: './select.page.html',
  styleUrls: ['./select.page.scss'],
})
export class SelectPage implements OnInit {
  tags$: Observable<Tag[]>
  tags: string[] = []
  constructor(
    private modalCtrl: ModalController,
    private navParams: NavParams,
  ) {
    this.tags$ = this.navParams.get('tags$')
    this.tags = this.navParams.get('tags')
  }

  ngOnInit() {
  }

  cancel() {
    this.modalCtrl.dismiss();
  }

  select() {
    this.modalCtrl.dismiss(this.tags);
  }

  change(tag: string) {
    if (this.has(tag)) {
      pull(this.tags, tag)
    } else {
      this.tags.push(tag)
    }
  }

  has(tag: string) {
    return includes(this.tags, tag)
  }
}
