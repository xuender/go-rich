import { Component } from '@angular/core';
import { ModalController, ActionSheetController } from '@ionic/angular';
import { Observable } from 'rxjs';

import { Tag } from './tag';
import { TagPage } from './tag/tag.page';
import { TagService } from './tag.service';
import { ObjsPage } from '../api/objs.page';

@Component({
  selector: 'app-tags',
  templateUrl: './tags.page.html',
  styleUrls: ['./tags.page.scss'],
})

export class TagsPage extends ObjsPage<Tag> {
  tags$: Observable<Tag[]>
  constructor(
    public tagService: TagService,
    protected modalCtrl: ModalController,
    protected actionSheetCtrl: ActionSheetController,
  ) {
    super(modalCtrl, actionSheetCtrl)
    this.tags$ = this.tagService.all$
  }

  get service() { return this.tagService }
  get page() { return TagPage }
  get newObj() { return {} }
  get title() { return '标签' }

  search(event: CustomEvent) {
    const txt = event.detail.value
    this.tags$ = this.tagService.search(txt)
  }
}
