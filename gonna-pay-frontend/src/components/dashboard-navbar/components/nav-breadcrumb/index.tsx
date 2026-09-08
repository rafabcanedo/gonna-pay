'use client'

import { Fragment } from 'react'
import { usePathname, useSearchParams } from 'next/navigation'
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from '@/components/ui/breadcrumb'
import { SEGMENT_LABELS } from './constants'

function buildBreadcrumbs(pathname: string, name: string | null): string[] {
  return pathname
    .split('/')
    .filter(Boolean)
    .filter((seg) => seg !== 'details')
    .map((seg) => SEGMENT_LABELS[seg] ?? name ?? seg)
}

export function NavBreadcrumb() {
  const pathname = usePathname()
  const searchParams = useSearchParams()
  const name = searchParams.get('name')

  const crumbs = buildBreadcrumbs(pathname, name)

  return (
    <Breadcrumb>
      <BreadcrumbList>
        {crumbs.map((crumb, index) => (
          <Fragment key={crumb}>
            {index > 0 && <BreadcrumbSeparator />}
            <BreadcrumbItem>
              {index === crumbs.length - 1 ? (
                <BreadcrumbPage>{crumb}</BreadcrumbPage>
              ) : (
                <span className="font-normal text-foreground">{crumb}</span>
              )}
            </BreadcrumbItem>
          </Fragment>
        ))}
      </BreadcrumbList>
    </Breadcrumb>
  )
}
