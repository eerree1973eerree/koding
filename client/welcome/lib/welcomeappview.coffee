kd = require 'kd'

module.exports = class WelcomeAppView extends kd.View

  constructor:->

    super

    @addSubView @welcome = new kd.CustomHTMLView
      tagName : 'section'
      partial : """
        <div class="artboard"></div>
        <h2>Welcome! let's get started.</h2>
        <p>
          Koding lets your team collaborate and
          work faster, with a development
          environment in the cloud
        </p>
        """
    { groupsController } = kd.singletons
    groupsController.ready =>
      currentGroup = groupsController.getCurrentGroup()
      currentGroup.fetchMyRoles (err, roles) =>
        return  kd.warn err  if err
        return  unless 'admin' in (roles ? [])

        @putInstructions()


  putInstructions: ->

      @welcome.setPartial """
        <ul>
          <li>
            <a href='/Admin/Stacks'>
              <cite>1</cite>
              <div>
                <span>Configure user environment</span>
                <span>Setup machine stacks & add their codebase</span>
              </div>
            </a>
          </li>
          <li>
            <a href='/Admin/Invitations'>
              <cite>2</cite>
              <div>
                <span>Invite your team</span>
                <span>Send out invites to your developers</span>
              </div>
            </a>
          </li>
        </ul>
        """