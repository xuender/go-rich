import { OnInit } from '@angular/core';
import { ModalController, ActionSheetController } from '@ionic/angular';
import { pull, includes } from 'lodash'

import { Obj } from './obj';
import { Paging } from './paging';
import { ObjsPage } from './objs.page';
import { TagService } from '../tags/tag.service';

export abstract class ObjsPlusPage<T extends Obj> extends ObjsPage<T> implements OnInit {
  tags: string[] = []
  paging: Paging<T> = { data: [], total: 0 }
  constructor(
    protected tagService: TagService,
    protected modalCtrl: ModalController,
    protected actionSheetCtrl: ActionSheetController,
  ) {
    super(modalCtrl, actionSheetCtrl)
  }

  ngOnInit() {
    this.reset()
  }

  abstract get tagKey(): string

  async reset() {
    console.log(this.service)
    this.paging = await this.service.paging$.toPromise()
  }

  async search(event: CustomEvent) {
    const txt = event.detail.value
    this.paging = await this.service.searchPaging(txt).toPromise()
  }

  async select() {
    const tags = await this.tagService.select(this.tagKey)
    if (tags) {
      this.tags = tags
      this.paging = await this.service.selectPaging(this.tags).toPromise()
    }
  }

  async load(event) {
    const p = await this.service.nextPaging$.toPromise()
    console.log('p', p)
    this.paging.data.push(...p.data)
    event.target.complete()
    event.target.disabled = p.total == p.data.length
  }

  hasTag(tag: string): boolean {
    return includes(this.tags, tag)
  }

  async removeTag(tag: string) {
    pull(this.tags, tag)
    this.paging = await this.service.paging$.toPromise()
  }
}
